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
	"errors"
	"github.com/jacci-ch/sdp-xlib/valuex"
	"strconv"
	"strings"
	"time"
)

// parseInt - parse a int value includes int8/int16/int/...
func parseInt[T valuex.IntType](src string) (T, error) {
	val, err := strconv.ParseInt(src, 10, 64)
	return PanicIf(T(val), src, err)
}

// parseUint - parse a uint value includes uint8/uint16/uint/...
func parseUint[T valuex.UintType](src string) (T, error) {
	val, err := strconv.ParseUint(src, 10, 64)
	return PanicIf(T(val), src, err)
}

// parseBool - parse a bool value.
func parseBool(src string) (bool, error) {
	val, err := strconv.ParseBool(src)
	return PanicIf(val, src, err)
}

// parseBool - parse a string value.
func parseStr(src string) (string, error) {
	return PanicIf(src, src, nil)
}

// parseBool - parse a []string value in "str1,str2,str3" format.
func parseStrArray(src string) (ret []string, err error) {
	for _, elem := range strings.Split(src, ",") {
		if val := strings.TrimSpace(elem); len(elem) != 0 {
			ret = append(ret, val)
		} else {
			return PanicIf([]string{}, src, errors.New("empty string array element"))
		}
	}

	return ret, nil
}

// parseBool - parse a time.Duration value.
func parseDuration(src string) (time.Duration, error) {
	duration, err := time.ParseDuration(src)
	return PanicIf(duration, src, err)
}

// parseAs - parse a value using given parse function and write to given dst address.
// This function will write the given default value to the dst address while error occurs.
func parseAs[T any](key string, dst *T, def T, parse func(string) (T, error)) error {
	if cfgVal, ok := gValues[key]; !ok || len(cfgVal) == 0 {
		*dst = def
		return nil
	} else if val, err := parse(cfgVal); err != nil {
		*dst = def
		return err
	} else {
		*dst = val
		return nil
	}
}
