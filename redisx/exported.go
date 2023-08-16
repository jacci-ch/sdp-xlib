// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package redisx

import (
	"context"
	"github.com/jacci-ch/sdp-xlib/jsonx"
	"github.com/jacci-ch/sdp-xlib/logx"
	"github.com/redis/go-redis/v9"
	"time"
)

// Provider
//
// A callback function definition for data provider. When we call
// GetXxxWith() functions we need to specify a provider function
// to generate data while the cached value is not found.
type Provider func() (any, error)

// Set
//
// Save key-value with given TTL to cache storage.
func Set(key string, value any, ttl time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), Cfg.WriteTimeoutDuration)
	defer cancel()

	return Client.Set(ctx, key, value, ttl).Err()
}

// Get
//
// Fetch value of given key as string type.
func Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Cfg.ReadTimeoutDuration)
	defer cancel()

	rsp, err := Client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		logx.Logger.Error(err)
	}

	return rsp, err
}

// GetWith
//
// Fetch value of given key as string value. If the given key is not existed,
// this function will call given provider function to generate the new data,
// and set the new generated data to cache storage with given TTL.
//
// This function will exit when the data provider returns an error.
func GetWith(key string, ttl time.Duration, provider Provider) (string, error) {
	rsp, err := Get(key)
	if err == nil || err != redis.Nil {
		return rsp, err
	}

	v, err := provider()
	if err != nil {
		logx.Logger.Error(err)
		return "", err
	}

	str, err := jsonx.MarshalToString(v)
	if err != nil {
		logx.Logger.Error(err)
		return "", err
	}

	if err = Set(key, str, ttl); err != nil {
		logx.Logger.Error(err)
	}

	return str, nil
}

func GetEx(key string, ttl time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Cfg.ReadTimeoutDuration)
	defer cancel()

	rsp, err := Client.GetEx(ctx, key, ttl).Result()
	if err != nil && err != redis.Nil {
		logx.Logger.Error(err)
	}

	return rsp, err
}

// GetAs
//
// Similar with Get but decode string value to given type.
// Argument v must be a value pointer of a specified type.
func GetAs(v any, key string) error {
	rsp, err := Get(key)
	if err != nil {
		return err
	}

	return jsonx.UnmarshalFromString(rsp, v)
}

// GetAsWith
//
// Similar with GetWith but decode string value to given type.
// Argument v must be a value pointer of a specified type.
func GetAsWith(v any, key string, ttl time.Duration, provider Provider) error {
	rsp, err := GetWith(key, ttl, provider)
	if err != nil {
		return err
	}

	return jsonx.UnmarshalFromString(rsp, v)
}

func GetAsEx(v any, key string, ttl time.Duration) error {
	rsp, err := GetEx(key, ttl)
	if err != nil {
		return err
	}

	return jsonx.Unmarshal([]byte(rsp), v)
}

// Del
//
// Delete a key from cache storage.
func Del(key string) {
	ctx, cancel := context.WithTimeout(context.Background(), Cfg.WriteTimeoutDuration)
	defer cancel()

	_, err := Client.Del(ctx, key).Result()
	if err != nil && err != redis.Nil {
		logx.Logger.Error(err)
	}
}
