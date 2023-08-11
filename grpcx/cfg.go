// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package grpcx

import (
	"errors"
	"fmt"
	"github.com/jacci-ch/sdp-xlib/cfgx"
	"github.com/jacci-ch/sdp-xlib/logx"
	"github.com/jacci-ch/sdp-xlib/valuex"
	"unsafe"
)

const (
	KeyAddr    = "server.rpc.listen.addr"
	KeyPort    = "server.rpc.listen.port"
	KeyNatAddr = "server.rpc.listen.nat.addr"
	KeyNatPort = "server.rpc.listen.nat.port"
)

var (
	DefCfg = &Config{Port: "9527"}
	Cfg    *Config
)

// Config
//
// A struct holds all configurations.
type Config struct {
	Addr string
	Port string

	NatAddr string
	NatPort string
}

// Validate
//
// Validates the configurations.
func (c *Config) Validate() error {
	if len(c.Port) == 0 {
		return errors.New("grpcx: listen port can't be none")
	}

	return nil
}

// ProbeNat
//
// Returns configured NAT address and port if not empty.
// Or returns local listened address and port.
func (c *Config) ProbeNat() (string, string) {
	return c.ProbeNatAddr(), c.ProbeNatPort()
}

// ProbeNatAddr
//
// Returns configured NAT address if not empty. Or returns
// local listened address.
func (c *Config) ProbeNatAddr() string {
	if len(c.NatAddr) != 0 {
		return c.NatAddr
	}

	return c.Addr
}

// ProbeNatPort
//
// Returns configured NAT port if not empty. Or returns
// local listened port.
func (c *Config) ProbeNatPort() string {
	if len(c.NatPort) != 0 {
		return c.NatPort
	}

	return c.Port
}

// Endpoint
//
// Returns the string value of address-port pair.
func (c *Config) Endpoint() string {
	return fmt.Sprintf("%v:%v", c.Addr, c.Port)
}

// ConfigKeys
//
// A struct holds all configuration keys.
type ConfigKeys struct {
	Addr    string
	Port    string
	NatAddr string
	NatPort string
}

// Validate
//
// Validates the keys.
func (c *ConfigKeys) Validate() error {
	if len(c.Addr) == 0 || len(c.Port) == 0 ||
		len(c.NatAddr) == 0 || len(c.NatPort) == 0 {
		return errors.New("grpcx: all keys can't be empty")
	}
	return nil
}

// LoadConfigsWith
//
// Load all configurations with given keys.
func LoadConfigsWith(keys *ConfigKeys) (*Config, error) {
	if keys == nil {
		return nil, errors.New("grpcx: arguments keys can't be nil")
	}

	if err := keys.Validate(); err != nil {
		return nil, err
	}

	return LoadConfigsWithGivenKeys(keys)
}

// LoadConfigsWithGivenKeys
//
// Load all configurations with given keys. This function do not
// validate the keys.
func LoadConfigsWithGivenKeys(keys *ConfigKeys) (*Config, error) {
	cfg := &Config{}

	if keys == nil {
		_ = cfgx.Def.ToStr(KeyAddr, &cfg.Addr, DefCfg.Addr)
		_ = cfgx.Def.ToStr(KeyPort, &cfg.Port, DefCfg.Port)
		_ = cfgx.Def.ToStr(KeyNatAddr, &cfg.NatAddr, DefCfg.NatAddr)
		_ = cfgx.Def.ToStr(KeyNatPort, &cfg.NatPort, DefCfg.NatPort)
	} else {
		_ = cfgx.Def.ToStr(keys.Addr, &cfg.Addr, DefCfg.Addr)
		_ = cfgx.Def.ToStr(keys.Port, &cfg.Port, DefCfg.Port)
		_ = cfgx.Def.ToStr(keys.NatAddr, &cfg.NatAddr, DefCfg.NatAddr)
		_ = cfgx.Def.ToStr(keys.NatPort, &cfg.NatPort, DefCfg.NatPort)
	}

	return cfg, nil
}

// LoadConfigs
//
// Load all configurations with default keys.
func LoadConfigs() (*Config, error) {
	return LoadConfigsWithGivenKeys(nil)
}

func init() {
	cfg, err := LoadConfigs()
	if err != nil {
		logx.Logger.Fatal(err)
	}

	valuex.SetPointer(unsafe.Pointer(&Cfg), unsafe.Pointer(cfg))
}
