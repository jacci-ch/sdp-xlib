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
	"github.com/jacci-ch/sdp-xlib/logx"
)

type OptionFunc func(s *Server)
type Provider func() GrpcServer

func OptKeys(keys *ConfigKeys) OptionFunc {
	return func(s *Server) {
		if cfg, err := loadConfigs(keys); err != nil {
			logx.Fatal(err)
		} else {
			s.cfg = cfg
		}
	}
}

func OptCfg(cfg *Config) OptionFunc {
	return func(s *Server) {
		s.cfg = cfg
	}
}

func OptServer(servers ...GrpcServer) OptionFunc {
	return func(s *Server) {
		s.servers = servers
	}
}

func OptProvider(providers ...Provider) OptionFunc {
	return func(s *Server) {
		var servers []GrpcServer
		for _, create := range providers {
			servers = append(servers, create())
		}

		s.servers = servers
	}
}
