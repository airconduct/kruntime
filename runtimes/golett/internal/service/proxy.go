package service

import (
	"fmt"

	"google.golang.org/grpc/resolver"

	"github.com/airconduct/kruntime/runtimes/golett/internal/store"
)

const (
	SCHEME_ACTOR = "actor"
)

func (s *GoletService) initResolver() {
	r := &pluginResolverBuilder{plugins: s.plugins}
	resolver.Register(r)
}

type pluginResolverBuilder struct {
	plugins store.Store[string, *pluginBundle]
}

func (b *pluginResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &pluginResolver{
		target:  target,
		cc:      cc,
		plugins: b.plugins,
	}
	r.ResolveNow(resolver.ResolveNowOptions{})
	return r, nil
}

func (b *pluginResolverBuilder) Scheme() string {
	return SCHEME_ACTOR
}

type pluginResolver struct {
	plugins store.Store[string, *pluginBundle]
	target  resolver.Target
	cc      resolver.ClientConn
}

func (r *pluginResolver) ResolveNow(resolver.ResolveNowOptions) {
	actorName := r.target.URL.Host
	id := r.target.URL.Query().Get("id")

	addrs := []resolver.Address{}
	r.plugins.Range(func(key string, value *pluginBundle) bool {
		if value.pid.Name != actorName {
			return true
		}
		if id != "" && value.pid.Id != id {
			return true
		}

		addr := value.client.ReattachConfig().Addr
		addrs = append(addrs, resolver.Address{
			Addr: fmt.Sprintf("%s://%s", addr.Network(), addr.String()),
		})
		return true
	})
	if err := r.cc.UpdateState(resolver.State{
		Addresses: addrs,
	}); err != nil {
		r.cc.ReportError(err)
	}
}

func (r *pluginResolver) Close() {
}

func ActorEndpoint(actorName, id string) string {
	if id == "" {
		return fmt.Sprintf("%s://%s", SCHEME_ACTOR, actorName)
	}
	return fmt.Sprintf("%s://%s?id=%s", SCHEME_ACTOR, actorName, id)
}
