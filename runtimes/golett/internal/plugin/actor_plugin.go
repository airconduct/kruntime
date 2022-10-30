package plugin

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/airconduct/kruntime/runtimes/golett/api"
	pb "github.com/airconduct/kruntime/runtimes/golett/api/proto"
)

type ActorPlugin struct {
	plugin.Plugin
	Impl api.Actor
}

var _ plugin.GRPCPlugin = &ActorPlugin{}

func (p *ActorPlugin) GRPCServer(b *plugin.GRPCBroker, s *grpc.Server) error {
	pb.RegisterActorServer(s, &ActorGRPCServer{actor: p.Impl})
	return nil
}

func (p *ActorPlugin) GRPCClient(ctx context.Context, b *plugin.GRPCBroker, cc *grpc.ClientConn) (interface{}, error) {
	return &ActorGRPCClient{ActorClient: pb.NewActorClient(cc)}, nil
}
