// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package redisx

import (
	"errors"
	"fmt"
	"github.com/jacci-ch/sdp-xlib/cfgx"
	"github.com/jacci-ch/sdp-xlib/timex"
	"time"
)

const (
	KeyAddrs    = "server.redis.address"
	KeyMaster   = "server.redis.master.name"
	KeyClient   = "server.redis.client.name"
	KeyUsername = "server.redis.username"
	KeyPassword = "server.redis.password"
	KeyDatabase = "server.redis.database"

	KeyReadTimeout  = "server.redis.read.timeout"
	KeyWriteTimeout = "server.redis.read.timeout"
)

var (
	Cfg    *Config
	DefCfg = &Config{
		ReadTimeout:  "30s",
		WriteTimeout: "30s",
	}
)

// Config
//
// A struct holds all configurations.
type Config struct {
	Addrs      []string
	MasterName string
	ClientName string
	Username   string
	Password   string
	Database   int

	ReadTimeout         string
	ReadTimeoutDuration time.Duration

	WriteTimeout         string
	WriteTimeoutDuration time.Duration
}

// Validate
//
// Validates the configurations.
func (c *Config) Validate() error {
	if len(c.Addrs) == 0 {
		return errors.New("redisx: addresses can't be empty")
	}

	tu, err := timex.ParseTimeUnit(c.ReadTimeout)
	if err != nil {
		return fmt.Errorf("redisx: invalid read timeout '%v'", err)
	}
	c.ReadTimeoutDuration = tu.Duration

	tu, err = timex.ParseTimeUnit(c.WriteTimeout)
	if err != nil {
		return fmt.Errorf("redisx: invalid write timeout '%v'", err)
	}
	c.WriteTimeoutDuration = tu.Duration

	return nil
}

// LoadConfigs
//
// Load all configurations from cfgx module.
func LoadConfigs() (*Config, error) {
	cfg := &Config{}

	_ = cfgx.Def.ToStrArray(KeyAddrs, &cfg.Addrs, DefCfg.Addrs)
	_ = cfgx.Def.ToStr(KeyMaster, &cfg.MasterName, DefCfg.MasterName)
	_ = cfgx.Def.ToStr(KeyClient, &cfg.ClientName, DefCfg.ClientName)
	_ = cfgx.Def.ToStr(KeyUsername, &cfg.Username, DefCfg.Username)
	_ = cfgx.Def.ToStr(KeyPassword, &cfg.Password, DefCfg.Password)
	_ = cfgx.Def.ToInt(KeyDatabase, &cfg.Database, DefCfg.Database)

	_ = cfgx.Def.ToStr(KeyReadTimeout, &cfg.ReadTimeout, DefCfg.ReadTimeout)
	_ = cfgx.Def.ToStr(KeyWriteTimeout, &cfg.WriteTimeout, DefCfg.WriteTimeout)

	return cfg, nil
}
