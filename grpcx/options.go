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
	"errors"

	"github.com/jacci-ch/sdp-xlib/logx"
)

type OptionFunc func(s *Server) error
type Provider func() (GrpcServer, error)

func OptKeys(keys *ConfigKeys) OptionFunc {
	return func(s *Server) error {
		if cfg, err := loadConfigs(keys); err != nil {
			logx.Fatal(err)
			return err
		} else {
			logx.Info("grpcx: use configuration loaded with given keys")
			s.cfg = cfg
			return nil
		}
	}
}

func OptCfg(cfg *Config) OptionFunc {
	return func(s *Server) error {
		if cfg == nil {
			err := errors.New("grpcx: argument cfg can't be nil")
			return logx.FatalErr(err)
		}

		logx.Info("grpcx: use specified configuration")
		s.cfg = cfg
		return nil
	}
}

func OptServer(servers ...GrpcServer) OptionFunc {
	return func(s *Server) error {
		if servers == nil || len(servers) == 0 {
			err := errors.New("grpcx: no rpc service provided")
			return logx.FatalErr(err)
		}

		s.servers = servers
		return nil
	}
}

func OptProvider(providers ...Provider) OptionFunc {
	return func(s *Server) error {
		if providers == nil || len(providers) == 0 {
			err := errors.New("grpcx: no rpc server providers specified")
			return logx.FatalErr(err)
		}

		var servers []GrpcServer
		for _, provider := range providers {
			if provider == nil {
				err := errors.New("grpcx: server provider can't be nil")
				return logx.FatalErr(err)
			}

			if server, err := provider(); err != nil {
				logx.Fatal(err)
				return err
			} else {
				servers = append(servers, server)
			}
		}

		s.servers = servers
		return nil
	}
}
