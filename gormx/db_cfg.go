package gormx

import (
	"errors"
	"fmt"
	"github.com/jacci-ch/sdp-xlib/cfgx"
)

const (
	KeyDefault = cfgx.Default

	KeyDsn      = "server.database.dsn"
	KeyUsername = "server.database.username"
	KeyPassword = "server.database.password"
	KeyHost     = "server.database.host"
	KeyPort     = "server.database.port"
	KeyDatabase = "server.database.database"
	KeyDebug    = "server.database.debug"
)

var (
	currCfg *Config
	defCfg  = &Config{
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

func (c *Config) Validate() error {
	if len(c.Dsn) != 0 {
		return nil
	}

	if len(c.Username) == 0 || len(c.Password) == 0 ||
		len(c.Host) == 0 || len(c.Port) == 0 ||
		len(c.Database) == 0 {
		return errors.New("gormx: invalid database configuration")
	}

	//<username>:<password>@tcp(<host>:<port>)/<database>
	c.Dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Username, c.Password, c.Host, c.Port, c.Database)
	return nil
}

func LoadConfigs() (*Config, error) {
	cfg := &Config{}

	_ = cfgx.ToStr(KeyDefault, KeyDsn, &cfg.Dsn, defCfg.Dsn)
	_ = cfgx.ToStr(KeyDefault, KeyUsername, &cfg.Username, defCfg.Username)
	_ = cfgx.ToStr(KeyDefault, KeyPassword, &cfg.Password, defCfg.Password)
	_ = cfgx.ToStr(KeyDefault, KeyHost, &cfg.Host, defCfg.Host)
	_ = cfgx.ToStr(KeyDefault, KeyPort, &cfg.Port, defCfg.Port)
	_ = cfgx.ToStr(KeyDefault, KeyDatabase, &cfg.Database, defCfg.Database)
	_ = cfgx.ToBool(KeyDefault, KeyDebug, &cfg.Debug, defCfg.Debug)

	return cfg, nil
}
