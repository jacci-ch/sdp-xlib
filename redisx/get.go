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
	"github.com/jacci-ch/sdp-xlib/jsonx"
	"github.com/jacci-ch/sdp-xlib/logx"
	"github.com/redis/go-redis/v9"
	"time"
)

// Get - retrieves the value of given key stored in redis server.
// This function call jsonx function to decode string values to given type.
func Get[T any](key string) (*T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Cfg.ReadTimeout)
	defer cancel()

	rsp, err := Client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		logx.Error(err)
	}

	var v T
	if err = jsonx.UnmarshalFromString(rsp, &v); err != nil {
		return &v, err
	}

	return nil, err
}

// GetEx - retrieves the value stored in redis server and prolong the
// TTL of the given key.
func GetEx[T any](key string, ttl time.Duration) (*T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Cfg.ReadTimeout)
	defer cancel()

	rsp, err := Client.GetEx(ctx, key, ttl).Result()
	if err != nil && err != redis.Nil {
		logx.Error(err)
	}

	var v T
	if err = jsonx.UnmarshalFromString(rsp, &v); err != nil {
		return &v, err
	}

	return &v, err
}

// GetWith - retrieves value of given key as given type. If the given key is not existed,
// this function will call given provider function to generate the new data,
// and set the new generated data to cache storage with given TTL.
//
// This function will exit when the data provider returns an error.
func GetWith[T any](key string, ttl time.Duration, provide Provider[T]) (*T, error) {
	rsp, err := Get[T](key)
	if err == nil || err != redis.Nil {
		return rsp, err
	}

	v, err := provide()
	if err != nil {
		logx.Error(err)
		return v, err
	}

	if str, err := jsonx.MarshalToString(v); err != nil {
		logx.Error(err)
		return nil, err
	} else if err = Set(key, str, ttl); err != nil {
		logx.Error(err)
	}

	return v, nil
}

// GetExWith - retrieves value of given key as given type. If the given
// key is not existed, this function will call given provider function to
// generate the new data, and set the new generated data to cache storage
// with given TTL.
//
// This function will exit when the data provider returns an error.
func GetExWith[T any](key string, ttl time.Duration, provide Provider[T]) (*T, error) {
	rsp, err := GetEx[T](key, ttl)
	if err == nil || err != redis.Nil {
		return rsp, err
	}

	v, err := provide()
	if err != nil {
		logx.Error(err)
		return v, err
	}

	if str, err := jsonx.MarshalToString(v); err != nil {
		logx.Error(err)
		return nil, err
	} else if err = Set(key, str, ttl); err != nil {
		logx.Error(err)
	}

	return v, nil
}

// Provider - a callback function definition for data provider.
// When we call GetWith() functions we need to specify a
// provider function to generate data while the cached value
// is not found.
type Provider[T any] func() (*T, error)
