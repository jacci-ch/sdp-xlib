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

package redisx

import (
	"context"
	"time"

	"github.com/jacci-ch/sdp-xlib/jsonx"
	"github.com/jacci-ch/sdp-xlib/logx"
	"github.com/redis/go-redis/v9"
)

func doGet[T any](get func(ctx context.Context) (string, error)) (*T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Cfg.ReadTimeout)
	defer cancel()

	if rsp, err := get(ctx); err == nil {
		if len(rsp) == 0 {
			return nil, nil
		} else {
			var v T
			if err = jsonx.UnmarshalFromString(rsp, &v); err != nil {
				return nil, err
			} else {
				return &v, err
			}
		}
	} else {
		if err != redis.Nil {
			logx.Error(err)
		}

		return nil, err
	}
}

// Get - retrieves the value of given key stored in redis server.
// This function call jsonx function to decode string values to given type.
func Get[T any](key string) (*T, error) {
	return doGet[T](func(ctx context.Context) (string, error) {
		return Client.Get(ctx, key).Result()
	})
}

// GetEx - retrieves the value stored in redis server and prolong the
// TTL of the given key.
func GetEx[T any](key string, ttl time.Duration) (*T, error) {
	return doGet[T](func(ctx context.Context) (string, error) {
		return Client.GetEx(ctx, key, ttl).Result()
	})
}

func doGetWithProvider[T any](key string, ttl time.Duration, get func() (*T, error), provider Provider[T]) (*T, error) {
	rsp, err := get()
	if err == nil || err != redis.Nil {
		return rsp, err
	}

	v, err := provider()
	if err != nil {
		logx.Error(err)
		return v, err
	}

	str := ""
	if v != nil {
		if str, err = jsonx.MarshalToString(v); err != nil {
			logx.Error(err)
			return nil, err
		}
	}

	if err = Set(key, str, ttl); err != nil {
		logx.Error(err)
	}

	return v, nil
}

// GetWithProvider - retrieves value of given key as given type. If the given key is not existed,
// this function will call given provider function to generate the new data,
// and set the new generated data to cache storage with given TTL.
//
// This function will exit when the data provider returns an error.
func GetWithProvider[T any](key string, ttl time.Duration, provider Provider[T]) (*T, error) {
	return doGetWithProvider(key, ttl, func() (*T, error) {
		return Get[T](key)
	}, provider)
}

// GetExWithProvider - retrieves value of given key as given type. If the given
// key is not existed, this function will call given provider function to
// generate the new data, and set the new generated data to cache storage
// with given TTL.
//
// This function will exit when the data provider returns an error.
func GetExWithProvider[T any](key string, ttl time.Duration, provider Provider[T]) (*T, error) {
	return doGetWithProvider(key, ttl, func() (*T, error) {
		return GetEx[T](key, ttl)
	}, provider)
}

// Provider - a callback function definition for data provider.
// When we call GetWithProvider() functions we need to specify a
// provider function to generate data while the cached value
// is not found.
type Provider[T any] func() (*T, error)
