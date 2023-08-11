// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package gormx

import (
	"errors"
	"fmt"
	"github.com/jacci-ch/sdp-xlib/cfgx"
)

const (
	KeyDsn      = "server.database.dsn"
	KeyUsername = "server.database.username"
	KeyPassword = "server.database.password"
	KeyHost     = "server.database.host"
	KeyPort     = "server.database.port"
	KeyDatabase = "server.database.database"
	KeyDebug    = "server.database.debug"
)

var (
	Cfg    *Config
	DefCfg = &Config{
		Port:  "3306",
		Debug: false,
	}
)

// Config
//
// A struct holds all configurations.
type Config struct {
	Dsn      string
	Username string
	Password string
	Host     string
	Port     string
	Database string
	Debug    bool
}

// Validate
//
// Validates the configurations.
func (c *Config) Validate() error {
	if len(c.Dsn) != 0 {
		return nil
	}

	if len(c.Username) == 0 || len(c.Password) == 0 ||
		len(c.Host) == 0 || len(c.Port) == 0 || len(c.Database) == 0 {
		return errors.New("gormx: invalid database configuration")
	}

	//<username>:<password>@tcp(<host>:<port>)/<database>
	c.Dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Username, c.Password, c.Host, c.Port, c.Database)
	return nil
}

// LoadConfigs
//
// Load all configurations from cfgx module.
func LoadConfigs() (*Config, error) {
	cfg := &Config{}

	_ = cfgx.Def.ToStr(KeyDsn, &cfg.Dsn, DefCfg.Dsn)
	_ = cfgx.Def.ToStr(KeyUsername, &cfg.Username, DefCfg.Username)
	_ = cfgx.Def.ToStr(KeyPassword, &cfg.Password, DefCfg.Password)
	_ = cfgx.Def.ToStr(KeyHost, &cfg.Host, DefCfg.Host)
	_ = cfgx.Def.ToStr(KeyPort, &cfg.Port, DefCfg.Port)
	_ = cfgx.Def.ToStr(KeyDatabase, &cfg.Database, DefCfg.Database)
	_ = cfgx.Def.ToBool(KeyDebug, &cfg.Debug, DefCfg.Debug)

	return cfg, nil
}
