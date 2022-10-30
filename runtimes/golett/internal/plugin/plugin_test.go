package plugin_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"testing"

	plugin "github.com/hashicorp/go-plugin"

	pb "github.com/airconduct/kruntime/runtimes/golett/api/proto"
	kplugin "github.com/airconduct/kruntime/runtimes/golett/internal/plugin"
	testapi "github.com/airconduct/kruntime/runtimes/golett/internal/plugin/test/api"
)

func TestPlugin(t *testing.T) {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "FOO",
			MagicCookieValue: "TEST",
		},
		Plugins: plugin.PluginSet{
			"actor": &kplugin.ActorPlugin{},
		},
		Cmd: exec.Command("./test/test"),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})
	rpcClient, err := client.Client()
	if err != nil {
		t.Error(err)
	}
	raw, err := rpcClient.Dispense("actor")
	if err != nil {
		t.Error(err)
	}
	actor := raw.(pb.ActorClient)
	payload := testapi.Task{Num: 2}
	body, _ := json.Marshal(payload)
	resp, err := actor.Call(context.Background(), &pb.RemoteMessage{
		Body:   body,
		Format: pb.DataFormat_JSON,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}
