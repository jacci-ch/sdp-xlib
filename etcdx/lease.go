// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package etcdx

import (
	"context"
	etcd "go.etcd.io/etcd/client/v3"
)

// Grant
//
// Call etcd function to generate a lease to use.
func Grant(ttl int64) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Cfg.WriteTimeoutDuration)
	defer cancel()

	rsp, err := gClient.Grant(ctx, ttl)
	if err != nil {
		return 0, err
	}

	return int64(rsp.ID), nil
}

// Revoke
//
// Revokes/Deletes the given lease with the default context and
// ignore the response.
func Revoke(lease int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), Cfg.WriteTimeoutDuration)
	defer cancel()

	_, err := gClient.Revoke(ctx, etcd.LeaseID(lease))
	return err
}

// KeepaliveOnce
//
// Send keepalive (heartbeat) message to etcd server only once.
func KeepaliveOnce(lease int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), Cfg.WriteTimeoutDuration)
	defer cancel()

	// ignore the result.
	_, err := gClient.KeepAliveOnce(ctx, etcd.LeaseID(lease))
	return err
}
