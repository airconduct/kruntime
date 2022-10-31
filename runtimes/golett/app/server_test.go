package app_test

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/airconduct/kruntime/runtimes/golett/api/proto"
)

func TestServer(t *testing.T) {
	cc, err := grpc.Dial("127.0.0.1:7770", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}

	client := pb.NewGoletClient(cc)
	resp, err := client.Load(context.Background(), &pb.LoadRequest{
		Name: "foo",
		Path: "../internal/plugin/test/test",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
