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
	"context"
	etcd "go.etcd.io/etcd/client/v3"
)

// putValue - calls gClient.Put to write KV to etcd server with background
// context and returns the response.
func putValue(key, value string, opts ...etcd.OpOption) (*etcd.PutResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Cfg.WriteTimeout)
	defer cancel()

	return gClient.Put(ctx, key, value, opts...)
}

func Put(key, value string) error {
	_, err := putValue(key, value)
	return err
}

// PutWithLease - writes a KV to etcd with given lease ID.
func PutWithLease(key, value string, lease int64) error {
	_, err := putValue(key, value, etcd.WithLease(etcd.LeaseID(lease)))
	return err
}

// PutWithNewLease - writes a KV to etcd with a new granted lease and returns the
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
