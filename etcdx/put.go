// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package etcdx

import (
	"context"
	etcd "go.etcd.io/etcd/client/v3"
)

// putValue
//
// Call gClient.Put to write KV to etcd server with background
// context and returns the response.
func putValue(key, value string, opts ...etcd.OpOption) (*etcd.PutResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Cfg.WriteTimeoutDuration)
	defer cancel()

	return gClient.Put(ctx, key, value, opts...)
}

func Put(key, value string) error {
	_, err := putValue(key, value)
	return err
}

// PutWithLease
//
// Write a KV to etcd with given lease ID.
func PutWithLease(key, value string, lease int64) error {
	_, err := putValue(key, value, etcd.WithLease(etcd.LeaseID(lease)))
	return err
}

// PutWithNewLease
//
// Write a KV to etcd with a new granted lease and returns the
// lease id to caller. The new generated lease will be revoked
// if errors occurred.
func PutWithNewLease(key, value string, ttl int64) (int64, error) {
	lease, err := Grant(ttl)
	if err != nil {
		return 0, err
	}

	if err = PutWithLease(key, value, lease); err != nil {
		_ = Revoke(lease)
		return 0, err
	}

	return lease, nil
}
