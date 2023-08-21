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

package etcdx

import (
	"errors"
	"github.com/jacci-ch/sdp-xlib/cfgx"
	"time"
)

const (
	KeyEndpoints    = "server.etcd.endpoints"
	KeyDialTimeout  = "server.etcd.dial.timeout"
	KeyReadTimeout  = "server.etcd.read.timeout"
	KeyWriteTimeout = "server.etcd.write.timeout"
)

var (
	Cfg    *Config
	DefCfg = &Config{
		DialTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
)

type Config struct {
	Endpoints []string

	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// validate - validates the configuration.
func (c *Config) validate() error {
	if len(c.Endpoints) == 0 {
		return errors.New("etcdx: endpoints no specified")
	}
	return nil
}

// loadConfigs - loads configurations from cfgx module.
func loadConfigs() *Config {
	cfg := &Config{}

	_ = cfgx.AsStrArray(KeyEndpoints, &cfg.Endpoints, DefCfg.Endpoints)
	_ = cfgx.AsDuration(KeyDialTimeout, &cfg.DialTimeout, DefCfg.DialTimeout)
	_ = cfgx.AsDuration(KeyReadTimeout, &cfg.ReadTimeout, DefCfg.ReadTimeout)
	_ = cfgx.AsDuration(KeyWriteTimeout, &cfg.WriteTimeout, DefCfg.WriteTimeout)

	return cfg
}
