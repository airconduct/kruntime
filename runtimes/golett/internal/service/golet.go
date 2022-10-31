package service

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
	"github.com/hashicorp/go-plugin"
	"golang.org/x/sys/unix"
	"google.golang.org/protobuf/proto"

	"github.com/airconduct/kruntime/runtimes/golett/api"
	pb "github.com/airconduct/kruntime/runtimes/golett/api/proto"
	kplugin "github.com/airconduct/kruntime/runtimes/golett/internal/plugin"
)

func New() *GoletService {
	return &GoletService{
		plugins: make(map[string]*plugin.Client),
		actors:  make(map[string]pb.ActorClient),
		pids:    make(map[string]*pb.PID),
	}
}

type GoletService struct {
	pb.UnimplementedGoletServer

	mutex   sync.RWMutex
	plugins map[string]*plugin.Client
	actors  map[string]pb.ActorClient
	pids    map[string]*pb.PID
}

var _ pb.GoletServer = &GoletService{}

func (s *GoletService) Invoke(ctx context.Context, in *pb.RemoteMessage) (out *pb.RemoteMessage, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	actorClient, ok := s.actors[in.To.Uuid]
	if !ok {
		return nil, fmt.Errorf("PID %s not exist", in.To.String())
	}

	return actorClient.Call(ctx, in)
}

func (s *GoletService) Load(ctx context.Context, in *pb.LoadRequest) (out *pb.LoadResponse, err error) {
	absPath, err := filepath.Abs(in.Path)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(absPath); err != nil {
		return nil, err
	} else if err := unix.Access(absPath, unix.X_OK); err != nil {
		return nil, fmt.Errorf("file %s is not executable, %v", absPath, err)
	}

	pid := &pb.PID{
		Uuid:         uuid.New().String(),
		Name:         in.Name,
		BinarySource: absPath,
		Host:         "",
		Endpoint:     "",
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()

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
	rpcClient, err := client.Client()
	if err != nil {
		return nil, err
	}
	raw, err := rpcClient.Dispense(api.PluginName)
	if err != nil {
		return nil, err
	}
	actorClient, ok := raw.(pb.ActorClient)
	if !ok {
		return nil, fmt.Errorf("failed to convert actor client")
	}

	pid.Endpoint = client.ReattachConfig().Addr.String()
	s.plugins[pid.Uuid] = client
	s.actors[pid.Uuid] = actorClient
	s.pids[pid.Uuid] = pid

	out = &pb.LoadResponse{Pid: pid}
	return
}

func (s *GoletService) Unload(ctx context.Context, in *pb.UnloadRequest) (out *pb.UnloadResponse, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	client, ok := s.plugins[in.Pid.Uuid]
	if !ok {
		return
	}
	if !client.Exited() {
		client.Kill()
	}

	delete(s.plugins, in.Pid.Uuid)
	delete(s.actors, in.Pid.Uuid)
	delete(s.pids, in.Pid.Uuid)

	return
}

func (s *GoletService) List(ctx context.Context, in *pb.ListRequest) (out *pb.ListResponse, err error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	out = &pb.ListResponse{}
	for i := range s.pids {
		out.Pids = append(out.Pids, proto.Clone(s.pids[i]).(*pb.PID))
	}
	return
}

func (s *GoletService) Shutdown() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	wg := sync.WaitGroup{}
	wg.Add(len(s.plugins))
	for k := range s.plugins {
		go func(client *plugin.Client) {
			defer wg.Done()
			client.Kill()
		}(s.plugins[k])
	}
	wg.Wait()
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
