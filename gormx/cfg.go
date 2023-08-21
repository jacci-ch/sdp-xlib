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

type Config struct {
	Dsn      string
	Username string
	Password string
	Host     string
	Port     string
	Database string
	Debug    bool
}

// validate - validates the configuration.
func (c *Config) validate() error {
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

// loadConfigs - loads and parses all configuration items from cfgx.
// This function will ignore errors.
func loadConfigs() *Config {
	cfg := &Config{}

	_ = cfgx.AsStr(KeyDsn, &cfg.Dsn, DefCfg.Dsn)
	_ = cfgx.AsStr(KeyUsername, &cfg.Username, DefCfg.Username)
	_ = cfgx.AsStr(KeyPassword, &cfg.Password, DefCfg.Password)
	_ = cfgx.AsStr(KeyHost, &cfg.Host, DefCfg.Host)
	_ = cfgx.AsStr(KeyPort, &cfg.Port, DefCfg.Port)
	_ = cfgx.AsStr(KeyDatabase, &cfg.Database, DefCfg.Database)
	_ = cfgx.AsBool(KeyDebug, &cfg.Debug, DefCfg.Debug)

	return cfg
}
