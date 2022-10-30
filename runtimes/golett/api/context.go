package api

import (
	"context"

	pb "github.com/airconduct/kruntime/runtimes/golett/api/proto"
)

type Context interface {
	context.Context

	Respond(resp interface{})
	RespondErr(err error)
	Message() Message
}

type Message interface {
	Unmarshal(v interface{}) error
	Raw() []byte
	Header() map[string]string
	From() *pb.PID
	To() *pb.PID
}
