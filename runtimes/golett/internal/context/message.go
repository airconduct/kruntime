package context

import (
	"encoding/json"
	"fmt"

	"github.com/airconduct/kruntime/runtimes/golett/api"
	pb "github.com/airconduct/kruntime/runtimes/golett/api/proto"
)

type genericMessage struct {
	data *pb.RemoteMessage
}

var _ api.Message = &genericMessage{}

func (msg *genericMessage) Unmarshal(v interface{}) error {
	switch msg.data.Format {
	case pb.DataFormat_JSON:
		return json.Unmarshal(msg.data.Body, v)
	}
	return fmt.Errorf("unknown data format %s", msg.data.Format.String())
}

func (msg *genericMessage) Raw() []byte {
	return msg.data.Body
}

func (msg *genericMessage) Header() map[string]string {
	return msg.data.Header
}

func (msg *genericMessage) From() *pb.PID {
	return msg.data.From
}

func (msg *genericMessage) To() *pb.PID {
	return msg.data.To
}
