// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package cfgv

import (
	"errors"
	"strconv"
	"strings"
)

var (
	Empty = errors.New("empty")
)

// Value
//
// A Value object is a representation of the value, a string value in fact,
// in configuration file.
type Value string

// Addr
// Returns the address pointer of the Value to avoid value-copy.
func (v Value) Addr() *Value {
	return &v
}

// ToStr
//
// Convert value to a string and store it to given address.
// This function returns an error if the value is empty.
func (v Value) ToStr(dst *string) error {
	if len(v) == 0 {
		return Empty
	}

	*dst = string(v)
	return nil
}

// ToStrArray
//
// Convert value to a string array. The format:
//
//	"str1,str2,str3,strN" => [
//		"str1",
//		"str2",
//		"str3",
//		"strN"
//	]
//
// Caution: we need a *[]string argument.
func (v Value) ToStrArray(dst *[]string) error {
	for _, src := range strings.Split(string(v), ",") {
		if src = strings.TrimSpace(src); len(src) != 0 {
			*dst = append(*dst, src)
		}
	}

	if len(*dst) == 0 {
		return Empty
	}

	return nil
}

// ToInt64
//
// Convert value to an int64 value.
func (v Value) ToInt64(dst *int64) error {
	if len(v) == 0 {
		return Empty
	}

	value, err := strconv.ParseInt(string(v), 10, 64)
	if err != nil {
		return err
	}

	*dst = value
	return nil
}

// ToInt32
//
// Convert value to an int32 value.
func (v Value) ToInt32(dst *int32) error {
	var value int64
	if err := v.ToInt64(&value); err != nil {
		return err
	}

	*dst = int32(value)
	return nil
}

// ToInt
//
// Convert value to an int value.
func (v Value) ToInt(dst *int) error {
	var value int64
	if err := v.ToInt64(&value); err != nil {
		return err
	}

	*dst = int(value)
	return nil
}

// ToBool
//
// Convert value to a bool value.
func (v Value) ToBool(dst *bool) error {
	if len(v) == 0 {
		return Empty
	}

	value, err := strconv.ParseBool(string(v))
	if err != nil {
		return err
	}

	*dst = value
	return nil
}
