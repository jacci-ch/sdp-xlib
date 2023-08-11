// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package grpcx

import (
	"google.golang.org/grpc"
)

// GrpcServer
//
// A grpc server implementation which implement this interface to
// register itself to grpc.Server.
type GrpcServer interface {
	Register(s *grpc.Server) error
}

// BeforeStartHook
//
// A hook execute before server start (before grpc.Server.Serve called).
type BeforeStartHook interface {
	BeforeServerStart(s *Server) error
}

// AfterStartHook
//
// A hook execute after server start (after grpc.Server.Serve called).
type AfterStartHook interface {
	AfterServerStart(s *Server) error
}

// BeforeStopHook
//
// A hook execute before server stop (after grpc.Server.stop called).
type BeforeStopHook interface {
	BeforeServerStop(s *Server)
}
