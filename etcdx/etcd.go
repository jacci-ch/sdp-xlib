// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package etcdx

import (
	"errors"
	"github.com/jacci-ch/sdp-xlib/logx"
	etcd "go.etcd.io/etcd/client/v3"
	"strings"
	"sync/atomic"
	"unsafe"
)

var (
	gClient *etcd.Client
)

// NewClient
//
// Generate a etcd client with given configuration.
func NewClient(cfg *Config) (*etcd.Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	etcdCfg := etcd.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: cfg.DialTimeoutDuration,
	}

	client, err := etcd.New(etcdCfg)
	if err != nil {
		return nil, errors.New("etcdx: " + err.Error())
	}

	return client, nil
}

// init
//
// Parses all configurations and generates a new etcd client with
// the configurations. This function store the etc client object
// in gClient value.
func init() {
	cfg, err := LoadConfigs()
	if err != nil {
		logx.Logger.Fatal(err)
		return
	}

	client, err := NewClient(cfg)
	if err != nil {
		logx.Logger.Fatal(err)
		return
	}

	// Atomically set the etcd client pointer and current in use configuration.
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&gClient)), unsafe.Pointer(client))
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&Cfg)), unsafe.Pointer(cfg))

	logx.Logger.Infof("etcdx: connect to %v", strings.Join(cfg.Endpoints, ","))
}
