// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package etcdx

import (
	"errors"
	"fmt"
	"github.com/jacci-ch/sdp-xlib/cfgx"
	"github.com/jacci-ch/sdp-xlib/timex"
	"time"
)

const (
	KeyEndpoints    = "server.etcd.endpoints"
	KeyDialTimeout  = "server.etcd.dial.timeout"
	KeyReadTimeout  = "server.etcd.read.timeout"
	KeyWriteTimeout = "server.etcd.write.timeout"
)

var (
	DefCfg = &Config{
		DialTimeout:  "10s",
		ReadTimeout:  "10s",
		WriteTimeout: "10s",
	}

	Cfg *Config
)

type Config struct {
	Endpoints []string

	DialTimeout         string
	DialTimeoutDuration time.Duration

	ReadTimeout         string
	ReadTimeoutDuration time.Duration

	WriteTimeout         string
	WriteTimeoutDuration time.Duration
}

// Validate
//
// Validates the configuration fields.
func (c *Config) Validate() error {
	if len(c.Endpoints) == 0 {
		return errors.New("etcdx: endpoint not configured")
	}

	tu, err := timex.ParseTimeUnit(c.DialTimeout)
	if err != nil {
		return fmt.Errorf("etcdx: invalid dial timeout: %v", err)
	}
	c.DialTimeoutDuration = tu.Duration

	tu, err = timex.ParseTimeUnit(c.ReadTimeout)
	if err != nil {
		return fmt.Errorf("etcdx: invalid read timeout: %v", err)
	}
	c.ReadTimeoutDuration = tu.Duration

	tu, err = timex.ParseTimeUnit(c.WriteTimeout)
	if err != nil {
		return fmt.Errorf("etcdx: invalid write timeout: %v", err)
	}
	c.WriteTimeoutDuration = tu.Duration

	return nil
}

// LoadConfigs
//
// Parses configurations from cfgx module.
func LoadConfigs() (*Config, error) {
	cfg := &Config{}

	_ = cfgx.Def.ToStrArray(KeyEndpoints, &cfg.Endpoints, DefCfg.Endpoints)
	_ = cfgx.Def.ToStr(KeyDialTimeout, &cfg.DialTimeout, DefCfg.DialTimeout)
	_ = cfgx.Def.ToStr(KeyReadTimeout, &cfg.ReadTimeout, DefCfg.ReadTimeout)
	_ = cfgx.Def.ToStr(KeyWriteTimeout, &cfg.WriteTimeout, DefCfg.WriteTimeout)

	return cfg, nil
}
