package grpcx

import (
	"fmt"
	"github.com/jacci-ch/sdp-xlib/grpcx/errors"
	"github.com/jacci-ch/sdp-xlib/logx"
	"net"
)

func NewListener(cfg *Config) (net.Listener, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	addr := fmt.Sprintf("%v:%v", cfg.ListenAddr, cfg.ListenPort)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, errors.Wrapper(err)
	}

	logx.Logger.Infof("grpcx: listen on: %v", addr)
	return listener, nil
}

func GetListener() (net.Listener, error) {
	if currCfg == nil {
		return nil, errors.ErrNoConfig
	}

	return NewListener(currCfg)
}
