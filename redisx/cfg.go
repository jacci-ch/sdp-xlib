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

package redisx

import (
	"errors"
	"time"

	"github.com/jacci-ch/sdp-xlib/cfgx"
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
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
)

// Config - a structure holds all configurations.
type Config struct {
	Addrs        []string
	MasterName   string
	ClientName   string
	Username     string
	Password     string
	Database     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// validate - validates the configurations.
func (c *Config) validate() error {
	if len(c.Addrs) == 0 {
		return errors.New("redisx: addresses can't be empty")
	}

	return nil
}

// loadConfigs - loads all configurations from cfgx module.
func loadConfigs() *Config {
	cfg := &Config{}

	_ = cfgx.AsStrArray(KeyAddrs, &cfg.Addrs, DefCfg.Addrs)
	_ = cfgx.AsStr(KeyMaster, &cfg.MasterName, DefCfg.MasterName)
	_ = cfgx.AsStr(KeyClient, &cfg.ClientName, DefCfg.ClientName)
	_ = cfgx.AsStr(KeyUsername, &cfg.Username, DefCfg.Username)
	_ = cfgx.AsStr(KeyPassword, &cfg.Password, DefCfg.Password)
	_ = cfgx.AsInt(KeyDatabase, &cfg.Database, DefCfg.Database)

	_ = cfgx.AsDuration(KeyReadTimeout, &cfg.ReadTimeout, DefCfg.ReadTimeout)
	_ = cfgx.AsDuration(KeyWriteTimeout, &cfg.WriteTimeout, DefCfg.WriteTimeout)

	return cfg
}
