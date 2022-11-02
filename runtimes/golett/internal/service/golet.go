package service

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"k8s.io/client-go/tools/cache"

	"github.com/airconduct/kruntime/runtimes/golett/api"
	pb "github.com/airconduct/kruntime/runtimes/golett/api/proto"
	kplugin "github.com/airconduct/kruntime/runtimes/golett/internal/plugin"
	"github.com/airconduct/kruntime/runtimes/golett/internal/store"
)

func New() *GoletService {
	svc := &GoletService{
		plugins: store.New[string, *pluginBundle](),
		conns:   store.New[string, *grpc.ClientConn](),
	}
	svc.initResolver()
	return svc
}

type GoletService struct {
	pb.UnimplementedGoletServer

	// actor id -> plugins
	cache.Indexer
	plugins store.Store[string, *pluginBundle]
	// actor endpoint -> conn
	conns store.Store[string, *grpc.ClientConn]
}

type pluginBundle struct {
	pid    *pb.PID
	client *plugin.Client
}

var _ pb.GoletServer = &GoletService{}

func (s *GoletService) Invoke(ctx context.Context, in *pb.RemoteMessage) (out *pb.RemoteMessage, err error) {
	actorClient, err := s.getActorClient(in.To)
	if err != nil {
		return nil, err
	}

	return actorClient.Call(ctx, in)
}

func (s *GoletService) Load(ctx context.Context, in *pb.LoadRequest) (out *pb.LoadResponse, err error) {
	// Check input valid
	absPath, err := filepath.Abs(in.Path)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(absPath); err != nil {
		return nil, err
	} else if err := unix.Access(absPath, unix.X_OK); err != nil {
		return nil, fmt.Errorf("file %s is not executable, %v", absPath, err)
	}
	// Generate PID
	pid := &pb.PID{
		Id:           uuid.New().String(),
		Name:         in.Name,
		BinarySource: absPath,
		Host:         "",
		Endpoint:     "",
	}
	// Create plugin client
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  api.ProtocolVersion,
			MagicCookieKey:   api.MagicCookieKey,
			MagicCookieValue: api.MagicCookieValue,
		},
		Plugins:          plugin.PluginSet{api.PluginName: &kplugin.ActorPlugin{}},
		Cmd:              exec.Command(absPath),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})
	// Start plugin
	addr, err := client.Start()
	if err != nil {
		return nil, err
	}

	pid.Endpoint = addr.String()
	s.plugins.Set(pid.Id, &pluginBundle{pid: pid, client: client})

	out = &pb.LoadResponse{Pid: pid}
	return
}

func (s *GoletService) Unload(ctx context.Context, in *pb.UnloadRequest) (out *pb.UnloadResponse, err error) {
	if in.Pid.Name == "" && in.Pid.Id == "" {
		return nil, fmt.Errorf("pid name or id is required")
	}

	ensureClientExited := func(cli *plugin.Client) {
		if !cli.Exited() {
			cli.Kill()
		}
	}
	// unload specified plugin
	if in.Pid.Id != "" {
		p, ok := s.plugins.Get(in.Pid.Id)
		if !ok {
			return nil, nil
		}
		ensureClientExited(p.client)
		s.plugins.Delete(in.Pid.Id)
		return &pb.UnloadResponse{}, nil
	}
	// unload plugins by its name
	s.plugins.Range(func(id string, p *pluginBundle) bool {
		if p.pid.Name == in.Pid.Name {
			s.plugins.Delete(id)
			s.conns.Delete(ActorEndpoint(p.pid.Name, ""))
			s.conns.Delete(ActorEndpoint(p.pid.Name, p.pid.Id))
		}
		return true
	})

	return
}

func (s *GoletService) List(ctx context.Context, in *pb.ListRequest) (out *pb.ListResponse, err error) {
	out = &pb.ListResponse{}
	s.plugins.Range(func(key string, value *pluginBundle) bool {
		out.Pids = append(out.Pids, proto.Clone(value.pid).(*pb.PID))
		return true
	})
	return
}

func (s *GoletService) Shutdown() {
	s.plugins.Range(func(key string, value *pluginBundle) bool {
		if !value.client.Exited() {
			value.client.Kill()
		}
		return true
	})
}

func (s *GoletService) getActorClient(pid *pb.PID) (actorClient pb.ActorClient, err error) {
	if pid.Name == "" {
		return nil, fmt.Errorf("actor name is empty")
	}

	target := ActorEndpoint(pid.Name, pid.Id)
	cc, ok := s.conns.Get(target)
	if !ok {
		cc, err = grpc.DialContext(context.Background(), target, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, err
		}
		s.conns.Set(target, cc)
	}
	return pb.NewActorClient(cc), nil
}

// func isExectuable(mode os.FileMode) bool {
// 	switch {
// 	case mode&0100 != 0:
// 		// executable by its owner
// 		return true
// 	case mode&0010 != 0:
// 		// executable by its group
// 		return true
// 	case mode&0001 != 0:
// 		// executable by others
// 		return true
// 	case mode&0111 != 0:
// 		// executable by any
// 		return true
// 	case mode&0111 == 0111:
// 		return true
// 	}
// 	return false
// }
