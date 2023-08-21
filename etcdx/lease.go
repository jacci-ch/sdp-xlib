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

// Grant - calls etcd function to generate a lease to use.
func Grant(ttl int64) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Cfg.WriteTimeout)
	defer cancel()

	rsp, err := gClient.Grant(ctx, ttl)
	if err != nil {
		return 0, err
	}

	return int64(rsp.ID), nil
}

// Revoke - revokes/deletes the given lease with the default context and
// ignore the response.
func Revoke(lease int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), Cfg.WriteTimeout)
	defer cancel()

	_, err := gClient.Revoke(ctx, etcd.LeaseID(lease))
	return err
}

// KeepaliveOnce - sends keepalive (heartbeat) message to etcd server only once.
func KeepaliveOnce(lease int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), Cfg.WriteTimeout)
	defer cancel()

	// ignore the result.
	_, err := gClient.KeepAliveOnce(ctx, etcd.LeaseID(lease))
	return err
}
