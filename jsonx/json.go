// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package jsonx

import (
	json "github.com/json-iterator/go"
)

// Marshal
//
// Same as json.Marshal.
func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

// MarshalIndent
//
// Same as json.MarshalIndent.
func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

// MarshalToString
//
// Same as Marshal but returns a string value.
func MarshalToString(v any) (string, error) {
	return json.MarshalToString(v)
}

// Unmarshal
//
// Same as json.Unmarshal.
func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

// UnmarshalFromString
//
// Same as Unmarshal but reads value from a string.
func UnmarshalFromString(str string, v any) error {
	return json.UnmarshalFromString(str, v)
}

// Encode
//
// Encode object into JSON string ignore all errors (use for debug).
func Encode(v any) string {
	if str, err := MarshalToString(v); err == nil {
		return str
	} else {
		panic(err)
	}
}
