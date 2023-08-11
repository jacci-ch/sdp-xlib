// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

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
	logx.Logger.Info("grpcx: register multicast resolver ...")
	resolver.Register(&Builder{})
}
