package api_test

import (
	"testing"

	"github.com/airconduct/kruntime/runtimes/golett/api/proto"
)

func TestProt(t *testing.T) {
	var _ proto.ActorClient
	var _ proto.ActorServer
}
