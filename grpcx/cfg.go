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
	"github.com/jacci-ch/sdp-xlib/cfgx"
)

var (
	DefKeys = &ConfigKeys{
		Addr:    "server.rpc.listen.addr",
		Port:    "server.rpc.listen.port",
		NatAddr: "server.rpc.listen.nat.addr",
		NatPort: "server.rpc.listen.nat.port",
	}
	Cfg    *Config
	DefCfg = &Config{Port: "9527"}
)

// Config - a struct holds all configurations.
type Config struct {
	Addr    string
	Port    string
	NatAddr string
	NatPort string
}

// validate - validates the configurations.
func (c *Config) validate() error {
	if len(c.Port) == 0 {
		return errors.New("grpcx: listen port can't be none")
	}

	return nil
}

// ProbeNat - retrieves configured NAT address and port if not empty.
// Or returns local listened address and port.
func (c *Config) ProbeNat() (string, string) {
	return c.ProbeNatAddr(), c.ProbeNatPort()
}

// ProbeNatAddr - retrieves configured NAT address if not empty.
// Or returns local listened address.
func (c *Config) ProbeNatAddr() string {
	if len(c.NatAddr) != 0 {
		return c.NatAddr
	}

	return c.Addr
}

// ProbeNatPort - retrieves configured NAT port if not empty.
// Or returns local listened port.
func (c *Config) ProbeNatPort() string {
	if len(c.NatPort) != 0 {
		return c.NatPort
	}

	return c.Port
}

// ConfigKeys - a struct holds all configuration keys.
type ConfigKeys struct {
	Addr    string
	Port    string
	NatAddr string
	NatPort string
}

// validate - validates the keys, all keys will not be empty.
func (c *ConfigKeys) validate() error {
	if len(c.Addr) == 0 || len(c.Port) == 0 ||
		len(c.NatAddr) == 0 || len(c.NatPort) == 0 {
		return errors.New("grpcx: all keys can't be empty")
	}
	return nil
}

func loadConfigs(keys *ConfigKeys) (*Config, error) {
	cfg := &Config{}

	_ = cfgx.AsStr(keys.Addr, &cfg.Addr, DefCfg.Addr)
	_ = cfgx.AsStr(keys.Port, &cfg.Port, DefCfg.Port)
	_ = cfgx.AsStr(keys.NatAddr, &cfg.NatAddr, DefCfg.NatAddr)
	_ = cfgx.AsStr(keys.NatPort, &cfg.NatPort, DefCfg.NatPort)

	return cfg, nil
}
