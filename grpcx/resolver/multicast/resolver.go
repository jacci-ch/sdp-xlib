package multicast

import (
	"fmt"
	"github.com/jacci-ch/sdp-xlib/logx"
	"google.golang.org/grpc/resolver"
	"strings"
)

const (
	Separator = ";"
)

type Resolver struct {
	target resolver.Target
	cc     resolver.ClientConn
}

func (r *Resolver) start() (resolver.Resolver, error) {
	var addrs []resolver.Address

	if list := strings.Split(r.target.URL.Host, Separator); len(list) > 0 {
		for _, addr := range list {
			addrs = append(addrs, resolver.Address{Addr: addr})
		}
	}

	if len(addrs) == 0 {
		return nil, fmt.Errorf("grpcx: invalid target url: %v", r.target.URL.String())
	}

	if err := r.cc.UpdateState(resolver.State{Addresses: addrs}); err != nil {
		return nil, fmt.Errorf("grpcx: %v", err)
	}

	return r, nil
}

func (r *Resolver) ResolveNow(o resolver.ResolveNowOptions) {
	logx.Logger.Warning("grpcx: ResolveNow() is not implemented in multicast resolver")
}

func (r *Resolver) Close() {}
