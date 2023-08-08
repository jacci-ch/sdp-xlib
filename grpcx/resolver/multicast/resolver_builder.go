package multicast

import (
	"github.com/jacci-ch/sdp-xlib/logx"
	"google.golang.org/grpc/resolver"
)

const (
	ResolverScheme = "multicast"
)

type ResolverBuilder struct{}

func (t *ResolverBuilder) Scheme() string {
	return ResolverScheme
}

func (t *ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	return (&Resolver{target: target, cc: cc}).start()
}

func init() {
	logx.Logger.Info("register multicast resolver ...")
	resolver.Register(&ResolverBuilder{})
}
