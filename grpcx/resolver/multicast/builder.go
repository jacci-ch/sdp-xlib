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
	"github.com/jacci-ch/sdp-xlib/logx"
	"google.golang.org/grpc/resolver"
)

const (
	Scheme = "multicast"
)

// Builder
//
// The builder of multicast resolver.
type Builder struct{}

// Scheme
//
// See description of the interface.
func (t *Builder) Scheme() string {
	return Scheme
}

// Build
//
// See description of the interface.
func (t *Builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	return (&Resolver{target: target, cc: cc}).start()
}

// init
//
// Register multicast resolver to grpc library.
func init() {
	logx.Info("grpcx: register multicast resolver ...")
	resolver.Register(&Builder{})
}
