// Copyright 2023 to now() The SDP Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package multicast

import (
	"fmt"
	"github.com/jacci-ch/sdp-xlib/logx"
	"google.golang.org/grpc/resolver"
	"strings"
)

const (
	Separator = ","
)

// Resolver
//
// A resolver to resolve multi targets connections.
type Resolver struct {
	target resolver.Target
	cc     resolver.ClientConn
}

// start
//
// Resolve the address from the dial target. For examples:
//
//	multicast://127.0.0.1:1234,127.0.0.1:2234
//
// will be parsed to string array.
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

	// update the resolver.ClientConn states with parsed addresses.
	if err := r.cc.UpdateState(resolver.State{Addresses: addrs}); err != nil {
		return nil, fmt.Errorf("grpcx: %v", err)
	}

	return r, nil
}

// ResolveNow
//
// See the interface description for more information.
func (r *Resolver) ResolveNow(_ resolver.ResolveNowOptions) {
	logx.Warn("grpcx: ResolveNow() is not implemented in multicast resolver")
}

// Close
//
// See the interface description for more information.
func (r *Resolver) Close() {}
