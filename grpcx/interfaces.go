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

package grpcx

import (
	"google.golang.org/grpc"
)

// GrpcServer - a grpc server implementation which implement this interface to
// register itself to grpc.Server.
type GrpcServer interface {
	RegisterRpc(s *grpc.Server) error
}

// BeforeStartHook
//
// A hook execute before server start (before grpc.Server.Serve called).
type BeforeStartHook interface {
	BeforeServerStart(cfg *Config, s *Server) error
}

// AfterStartHook
//
// A hook execute after server start (after grpc.Server.Serve called).
type AfterStartHook interface {
	AfterServerStart(cfg *Config, s *Server) error
}

// BeforeStopHook
//
// A hook execute before server stop (after grpc.Server.stop called).
type BeforeStopHook interface {
	BeforeServerStop(cfg *Config, s *Server)
}
