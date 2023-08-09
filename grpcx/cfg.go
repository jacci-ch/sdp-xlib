package grpcx

import (
	"github.com/jacci-ch/sdp-xlib/cfgx"
	"github.com/jacci-ch/sdp-xlib/grpcx/errors"
	"github.com/jacci-ch/sdp-xlib/logx"
	"github.com/jacci-ch/sdp-xlib/valuex"
	"unsafe"
)

const (
	KeyListenAddr = "server.rpc.listen.addr"
	KeyListenPort = "server.rpc.listen.port"
	KeyNatAddr    = "server.rpc.listen.nat.addr"
	KeyNatPort    = "server.rpc.listen.nat.port"
)

var (
	defCfg  = &Config{ListenPort: "9527"}
	currCfg *Config
)

type Config struct {
	ListenAddr string
	ListenPort string

	NatAddr string
	NatPort string
}

func (c *Config) Validate() error {
	if len(c.ListenPort) == 0 {
		return errors.ErrInvalidConfig
	}

	return nil
}

type ConfigKeys struct {
	ListenAddr string
	ListenPort string
	NatAddr    string
	NatPort    string
}

func LoadConfigs(keys *ConfigKeys) (*Config, error) {
	if keys == nil {
		return nil, errors.ErrInvalidArgs
	}

	return loadConfigsByKeys(keys)
}

func loadConfigsByKeys(keys *ConfigKeys) (*Config, error) {
	cfg := &Config{}

	if keys == nil {
		_ = cfgx.Def.ToStr(KeyListenAddr, &cfg.ListenAddr, defCfg.ListenAddr)
		_ = cfgx.Def.ToStr(KeyListenPort, &cfg.ListenPort, defCfg.ListenPort)
		_ = cfgx.Def.ToStr(KeyNatAddr, &cfg.NatAddr, defCfg.NatAddr)
		_ = cfgx.Def.ToStr(KeyNatPort, &cfg.NatPort, defCfg.NatPort)
	} else {
		_ = cfgx.Def.ToStr(keys.ListenAddr, &cfg.ListenAddr, defCfg.ListenAddr)
		_ = cfgx.Def.ToStr(keys.ListenPort, &cfg.ListenPort, defCfg.ListenPort)
		_ = cfgx.Def.ToStr(keys.NatAddr, &cfg.NatAddr, defCfg.NatAddr)
		_ = cfgx.Def.ToStr(keys.NatPort, &cfg.NatPort, defCfg.NatPort)
	}

	return cfg, nil
}

func loadConfigs() (*Config, error) {
	return loadConfigsByKeys(nil)
}

func init() {
	cfg, err := loadConfigs()
	if err != nil {
		logx.Logger.Fatal(err)
	}

	valuex.SetPointer(unsafe.Pointer(&currCfg), unsafe.Pointer(cfg))
}
