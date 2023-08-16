// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package etcdx

// KV
//
// A struct like etcd.KV struct but store the key/value
// in string type instead of []byte.
type KV struct {
	Key   string
	Value string
}