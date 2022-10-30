package main

import (
	plugin "github.com/hashicorp/go-plugin"

	"github.com/airconduct/kruntime/runtimes/golett/api"
	kplugin "github.com/airconduct/kruntime/runtimes/golett/internal/plugin"
	testapi "github.com/airconduct/kruntime/runtimes/golett/internal/plugin/test/api"
)

type actor struct{}

func (a *actor) Receive(ctx api.Context) {
	t := &testapi.Task{}
	err := ctx.Message().Unmarshal(t)
	if err != nil {
		ctx.RespondErr(err)
		return
	}
	ctx.Respond(t.Num + 1)
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "FOO",
			MagicCookieValue: "TEST",
		},
		Plugins: plugin.PluginSet{
			"actor": &kplugin.ActorPlugin{Impl: &actor{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
