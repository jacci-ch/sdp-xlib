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

package cfgx

import (
	"github.com/jacci-ch/sdp-xlib/valuex"
	"time"
)

var (
	gValues map[string]string
)

func AsInt[T valuex.IntType](key string, dst *T, def T) error {
	return parseAs(key, dst, def, parseInt[T])
}

func AsUint[T valuex.UintType](key string, dst *T, def T) error {
	return parseAs(key, dst, def, parseUint[T])
}

func AsStr(key string, dst *string, def string) error {
	return parseAs(key, dst, def, parseStr)
}

func AsBool(key string, dst *bool, def bool) error {
	return parseAs(key, dst, def, parseBool)
}

func AsStrArray(key string, dst *[]string, def []string) error {
	return parseAs(key, dst, def, parseStrArray)
}

func AsDuration(key string, dst *time.Duration, def time.Duration) error {
	return parseAs(key, dst, def, parseDuration)
}
