package service_test

import (
	"context"
	"fmt"
	"net"
	"os"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/airconduct/kruntime/runtimes/golett/api/proto"
	testapi "github.com/airconduct/kruntime/runtimes/golett/internal/plugin/test/api"
	"github.com/airconduct/kruntime/runtimes/golett/internal/service"
)

func TestGolet(t *testing.T) {
	s := grpc.NewServer()

	gs := service.New()
	pb.RegisterGoletServer(s, gs)

	lis, err := serverListener_unix()
	if err != nil {
		t.Fatal(err)
	}
	go s.Serve(lis)
	defer func() {
		gs.Shutdown()
		s.Stop()
	}()

	cc, err := grpc.Dial(fmt.Sprintf("%s://%s", lis.Addr().Network(), lis.Addr().String()), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}

	client := pb.NewGoletClient(cc)

	_, err = client.Load(context.Background(), &pb.LoadRequest{
		Name: "foo", Path: "../plugin/test/test",
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Load(context.Background(), &pb.LoadRequest{
		Name: "foo", Path: "../plugin/test/test",
	})
	if err != nil {
		t.Fatal(err)
	}

	task := &testapi.Task{Num: 10}
	out, err := client.Invoke(context.Background(), &pb.RemoteMessage{
		To:     &pb.PID{Name: "foo"},
		Format: pb.DataFormat_JSON,
		Body:   task.Unmarshal(),
	})
	if err != nil {
		t.Fatal(err)
	}
	if string(out.Body) != "11" {
		t.Error("out should be 11")
	}
}

func serverListener_unix() (net.Listener, error) {
	tf, err := os.CreateTemp("", "test")
	if err != nil {
		return nil, err
	}
	path := tf.Name()

	// Close the file and remove it because it has to not exist for
	// the domain socket.
	if err := tf.Close(); err != nil {
		return nil, err
	}
	if err := os.Remove(path); err != nil {
		return nil, err
	}

	l, err := net.Listen("unix", path)
	if err != nil {
		return nil, err
	}

	// Wrap the listener in rmListener so that the Unix domain socket file
	// is removed on close.
	return &rmListener{
		Listener: l,
		Path:     path,
	}, nil
}

// rmListener is an implementation of net.Listener that forwards most
// calls to the listener but also removes a file as part of the close. We
// use this to cleanup the unix domain socket on close.
type rmListener struct {
	net.Listener
	Path string
}

func (l *rmListener) Close() error {
	// Close the listener itself
	if err := l.Listener.Close(); err != nil {
		return err
	}

	// Remove the file
	return os.Remove(l.Path)
}
