package context

import (
	"context"
	"encoding/json"

	"github.com/airconduct/kruntime/runtimes/golett/api"
	pb "github.com/airconduct/kruntime/runtimes/golett/api/proto"
)

func WithRemote(parent context.Context, msg *pb.RemoteMessage) *remoteContext {
	ctx := &remoteContext{
		Context: parent, msg: &genericMessage{data: msg},
		from: msg.From,
		to:   msg.To,
	}
	return ctx
}

type remoteContext struct {
	context.Context

	from, to *pb.PID
	resp     *pb.RemoteMessage
	msg      api.Message
	waitErr  error
}

var _ api.Context = &remoteContext{}

func (ctx *remoteContext) Respond(resp interface{}) {
	msg := &pb.RemoteMessage{
		Header: map[string]string{},
		From:   ctx.to,
		To:     ctx.from,
		Format: pb.DataFormat_JSON,
	}

	raw, err := json.Marshal(resp)
	if err != nil {
		ctx.waitErr = err
		return
	}
	msg.Body = raw
	ctx.resp = msg
}

func (ctx *remoteContext) RespondErr(err error) {
	ctx.waitErr = err
}

func (ctx *remoteContext) Message() api.Message {
	return ctx.msg
}

func (ctx *remoteContext) Response() *pb.RemoteMessage {
	return ctx.resp
}

// Context.Err implementation
func (ctx *remoteContext) Err() error {
	if ctx.waitErr == nil {
		return ctx.Context.Err()
	}
	return ctx.waitErr
}
