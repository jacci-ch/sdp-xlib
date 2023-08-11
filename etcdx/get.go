// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package etcdx

import (
	"context"
	"github.com/jacci-ch/sdp-xlib/jsonx"
	etcd "go.etcd.io/etcd/client/v3"
)

// getRsp
//
// Execute gClient.Get() with background context and returns
// the response object.
func getRsp(key string, opts ...etcd.OpOption) (*etcd.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Cfg.ReadTimeoutDuration)
	defer cancel()

	return gClient.Get(ctx, key, opts...)
}

// GetBytes
//
// Returns the value of given key in []byte type.
func GetBytes(key string) ([]byte, error) {
	rsp, err := getRsp(key)
	if err == nil && rsp.Count > 0 && len(rsp.Kvs) != 0 {
		return rsp.Kvs[0].Value, nil
	}
	return nil, err
}

func Get(key string) (string, error) {
	bytes, err := GetBytes(key)
	if err != nil {
		return "", nil
	}

	return string(bytes), nil
}

// GetAs
//
// Returns the value of given key in given type. This function use
// jsonx.Unmarshal function to parse bytes into given i address.
func GetAs(key string, i any) error {
	bytes, err := GetBytes(key)
	if err != nil {
		return err
	}

	return jsonx.Unmarshal(bytes, i)
}

// GetKeys
//
// Returns all keys with given prefix in etcd clusters. Caution:
// a very short prefix can match a large mount of keys. Please use
// this function carefully.
func GetKeys(prefix string) ([]string, error) {
	rsp, err := getRsp(prefix, etcd.WithPrefix(), etcd.WithKeysOnly())
	if err != nil {
		return nil, err
	}

	result := make([]string, len(rsp.Kvs))
	for cc, kv := range rsp.Kvs {
		result[cc] = string(kv.Key)
	}

	return result, nil
}

// GetByPrefix
//
// Returns all KVs matches given prefix and returns an array.
func GetByPrefix(prefix string) ([]*KV, error) {
	rsp, err := getRsp(prefix, etcd.WithPrefix())
	if err != nil {
		return nil, err
	}

	result := make([]*KV, len(rsp.Kvs))
	for cc, item := range rsp.Kvs {
		result[cc] = &KV{Key: string(item.Key), Value: string(item.Value)}
	}

	return result, nil
}

// GetByPrefixAsMap
//
// Returns all KVs matches given prefix and converts to a map[<key>]<value>.
func GetByPrefixAsMap(prefix string) (map[string]string, error) {
	rsp, err := getRsp(prefix, etcd.WithPrefix())
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, item := range rsp.Kvs {
		result[string(item.Key)] = string(item.Value)
	}

	return result, nil
}
