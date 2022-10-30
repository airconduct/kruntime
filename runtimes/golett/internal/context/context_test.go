package context_test

import (
	"context"
	"testing"

	pb "github.com/airconduct/kruntime/runtimes/golett/api/proto"
	kcontext "github.com/airconduct/kruntime/runtimes/golett/internal/context"
)

func TestContext(t *testing.T) {
	ctx := kcontext.WithRemote(context.Background(), &pb.RemoteMessage{
		From:   &pb.PID{Name: "test", Uuid: "1"},
		To:     &pb.PID{Name: "test", Uuid: "2"},
		Body:   []byte("{\"test\":1}"),
		Format: pb.DataFormat_JSON,
	})
	v := map[string]int{}
	err := ctx.Message().Unmarshal(&v)
	if err != nil {
		t.Error(err)
	}
	if v["test"] != 1 {
		t.Error("failed")
	}
	ctx.Respond(map[string]string{"foo": "test"})
	resp := ctx.Response()
	if string(resp.Body) != "{\"foo\":\"test\"}" {
		t.Error(string(resp.Body))
	}
}
