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

package jsonx

import (
	json "github.com/json-iterator/go"
)

// Marshal - same as json.Marshal.
func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

// MarshalIndent - same as json.MarshalIndent.
func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

// MarshalToString - same as Marshal but returns a string value.
func MarshalToString(v any) (string, error) {
	return json.MarshalToString(v)
}

// Unmarshal - same as json.Unmarshal.
func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

// UnmarshalFromString - same as Unmarshal but reads value from a string.
func UnmarshalFromString(str string, v any) error {
	return json.UnmarshalFromString(str, v)
}

// Decode - decodes given string to a specified value.
func Decode[T any](src string) (v T, err error) {
	err = json.UnmarshalFromString(src, &v)
	return v, err
}

// Encode - encodes given value to string value.
func Encode[T any](v T) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}
