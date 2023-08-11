// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package valuex

import (
	"errors"
	"strconv"
	"strings"
)

// NumUnit
//
// A structure holds num-unit value. Examples:
//
//	3s - 3 seconds
//	4m - 4 miles
//
// So we can use readable string in our configuration files.
// NumUnit do not explain the unit, for example:
// 4m can be 4 miles, 4 minutes or what ever things.
type NumUnit struct {
	Value int
	Unit  string
}

// ParseNumUnit
//
// Parse str value into NumUnit value. This function returns an error
// while the str value is not a valid num-unit string.
func ParseNumUnit(str string) (*NumUnit, error) {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return nil, errors.New("valuex: argument str can't be nil")
	}

	index := FirstNonDigit(str)
	v, err := strconv.ParseInt(str[0:index], 10, 64)
	if err != nil {
		return nil, err
	}

	return &NumUnit{Value: int(v), Unit: strings.TrimSpace(str[index:])}, nil
}
