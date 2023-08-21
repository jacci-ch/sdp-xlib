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
	"github.com/jacci-ch/sdp-xlib/jsonx"
	etcd "go.etcd.io/etcd/client/v3"
)

// getRsp - executes gClient.Get() with background context and retrieves
// the response object.
func getRsp(key string, opts ...etcd.OpOption) (*etcd.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Cfg.ReadTimeout)
	defer cancel()

	return gClient.Get(ctx, key, opts...)
}

// GetBytes - retrieves the value of given key in []byte type.
func GetBytes(key string) ([]byte, error) {
	rsp, err := getRsp(key)
	if err == nil && rsp.Count > 0 && len(rsp.Kvs) != 0 {
		return rsp.Kvs[0].Value, nil
	}
	return nil, err
}

func Get[T any](key string) (ret T, _ error) {
	bytes, err := GetBytes(key)
	if err == nil && len(bytes) != 0 {
		err = jsonx.Unmarshal(bytes, &ret)
		return ret, err
	}

	return ret, nil
}

// GetKeys - retrieves all keys with given prefix in etcd clusters. Caution:
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

// GetAll - retrieves all KVs matches given prefix and returns an array.
func GetAll[T any](prefix string) ([]*KV[T], error) {
	rsp, err := getRsp(prefix, etcd.WithPrefix())
	if err != nil {
		return nil, err
	}

	result := make([]*KV[T], len(rsp.Kvs))
	for cc, item := range rsp.Kvs {
		var v T
		if err = jsonx.Unmarshal(item.Value, &v); err != nil {
			return nil, err
		}
		result[cc] = &KV[T]{Key: string(item.Key), Value: &v}
	}

	return result, nil
}

// GetAllAsMap - retrieves all KVs matches given prefix and converts to a map[<key>]<value>.
func GetAllAsMap[T any](prefix string) (map[string]*T, error) {
	rsp, err := getRsp(prefix, etcd.WithPrefix())
	if err != nil {
		return nil, err
	}

	result := make(map[string]*T)
	for _, item := range rsp.Kvs {
		var v T
		if err = jsonx.Unmarshal(item.Value, &v); err != nil {
			return nil, err
		}
		result[string(item.Key)] = &v
	}

	return result, nil
}
