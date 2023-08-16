// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package redisx

import (
	"context"
	"github.com/jacci-ch/sdp-xlib/logx"
	"github.com/redis/go-redis/v9"
	"strings"
	"sync"
)

var (
	Client redis.UniversalClient
	gLock  = sync.Mutex{}
)

// NewClient
//
// Create a new redis Universal Client. See redis documents for more
// about redis.UniversalClient.
func NewClient(cfg *Config) redis.UniversalClient {
	return redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      cfg.Addrs,
		Username:   cfg.Username,
		Password:   cfg.Password,
		ClientName: cfg.ClientName,
		MasterName: cfg.MasterName,
		DB:         cfg.Database,
	})
}

// GetClient
//
// Returns a redis UniversalClient object.
func GetClient() redis.UniversalClient {
	if Client == nil {
		gLock.Lock()
		if Client == nil {
			Client = NewClient(Cfg)
		}
		gLock.Unlock()
	}

	// Test whether redis client is available.
	if err := Client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	logx.Logger.Infof("redisx: create redis client with address %v", strings.Join(Cfg.Addrs, ","))
	return Client
}
