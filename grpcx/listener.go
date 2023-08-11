// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package grpcx

import (
	"errors"
	"net"
)

// NewListenerWithCfg
//
// Create a listener with given configurations.
func NewListenerWithCfg(cfg *Config) (net.Listener, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	listener, err := net.Listen("tcp", cfg.Endpoint())
	if err != nil {
		return nil, errors.New("grpcx: " + err.Error())
	}

	return listener, nil
}

// NewListenerWithKeys
//
// Create a listener with given configuration keys. This function
// loads all configurations with given keys to generate a Config
// object and then calls to NewListenerWithCfg function.
func NewListenerWithKeys(keys *ConfigKeys) (net.Listener, error) {
	cfg, err := LoadConfigsWith(keys)
	if err != nil {
		return nil, err
	}

	return NewListenerWithCfg(cfg)
}

// NewListener
//
// Create a listener with default configurations.
func NewListener() (net.Listener, error) {
	return NewListenerWithCfg(Cfg)
}
