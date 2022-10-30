package plugin

import (
	"context"

	"github.com/airconduct/kruntime/runtimes/golett/api"
	pb "github.com/airconduct/kruntime/runtimes/golett/api/proto"
	kcontext "github.com/airconduct/kruntime/runtimes/golett/internal/context"
)

type ActorGRPCServer struct {
	pb.UnimplementedActorServer

	actor api.Actor
}

var _ pb.ActorServer = &ActorGRPCServer{}

func (s *ActorGRPCServer) Call(ctx context.Context, req *pb.RemoteMessage) (resp *pb.RemoteMessage, err error) {
	actorCtx := kcontext.WithRemote(ctx, req)
	s.actor.Receive(actorCtx)
	return actorCtx.Response(), actorCtx.Err()
}
