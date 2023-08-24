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

package grpcx

import (
	"errors"
	"fmt"
	"net"
)

// NewListenerWithCfg - creates a listener with given configurations.
func NewListenerWithCfg(cfg *Config) (net.Listener, error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port))
	if err != nil {
		return nil, errors.New("grpcx: " + err.Error())
	}

	return listener, nil
}

// NewListenerWithKeys - create a listener with given configuration keys.
// This function loads all configurations with given keys to generate a Config
// object and then calls to NewListenerWithCfg function.
func NewListenerWithKeys(keys *ConfigKeys) (net.Listener, error) {
	if cfg, err := loadConfigs(keys); err != nil {
		return nil, err
	} else {
		return NewListenerWithCfg(cfg)
	}
}

// NewListener - creates a listener with default configurations.
func NewListener() (net.Listener, error) {
	return NewListenerWithCfg(gCfg)
}
