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

package etcdx

import (
	"errors"
	etcd "go.etcd.io/etcd/client/v3"
)

var (
	gClient *etcd.Client
)

// KV - a struct like etcd.KV struct but store the key/value
// in string type instead of []byte.
type KV[T any] struct {
	Key   string
	Value *T
}

// NewClient - creates an etcd client with given configuration.
func NewClient(cfg *Config) (*etcd.Client, error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	etcdCfg := etcd.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: cfg.DialTimeout,
	}

	client, err := etcd.New(etcdCfg)
	if err != nil {
		return nil, errors.New("etcdx: " + err.Error())
	}

	return client, nil
}
