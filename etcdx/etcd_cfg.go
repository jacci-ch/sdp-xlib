package etcdx

import (
	"errors"
	"fmt"
	"github.com/jacci-ch/sdp-xlib/cfgx"
	"github.com/jacci-ch/sdp-xlib/timex"
	"time"
)

const (
	KeyDefault = cfgx.Default

	KeyEndpoints    = "server.etcd.endpoints"
	KeyDialTimeout  = "server.etcd.dial.timeout"
	KeyReadTimeout  = "server.etcd.read.timeout"
	KeyWriteTimeout = "server.etcd.write.timeout"
)

var (
	defCfg = &Config{
		DialTimeout:  "30s",
		ReadTimeout:  "30s",
		WriteTimeout: "30s",
	}

	// the current in use configurations.
	currCfg *Config
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
// Parses the configuration fields from cfgx module.
func LoadConfigs() (*Config, error) {
	cfg := &Config{}

	_ = cfgx.ToStrArray(KeyDefault, KeyEndpoints, &cfg.Endpoints, defCfg.Endpoints)
	_ = cfgx.ToStr(KeyDefault, KeyDialTimeout, &cfg.DialTimeout, defCfg.DialTimeout)
	_ = cfgx.ToStr(KeyDefault, KeyReadTimeout, &cfg.ReadTimeout, defCfg.ReadTimeout)
	_ = cfgx.ToStr(KeyDefault, KeyWriteTimeout, &cfg.WriteTimeout, defCfg.WriteTimeout)

	return cfg, nil
}
