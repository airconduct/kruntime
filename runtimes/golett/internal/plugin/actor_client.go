package plugin

import (
	pb "github.com/airconduct/kruntime/runtimes/golett/api/proto"
)

type ActorGRPCClient struct {
	pb.ActorClient
}
