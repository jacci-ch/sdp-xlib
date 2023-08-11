// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package timex

import (
	"fmt"
	"github.com/jacci-ch/sdp-xlib/valuex"
	"time"
)

var (
	gDurationMap = map[string]time.Duration{
		"ms": time.Millisecond,
		"m":  time.Minute,
		"s":  time.Second,
		"h":  time.Hour,
	}
)

// TimeUnit
//
// A struct holds time-with-unit string. Examples:
//
//	1h -  1  * time.Hour
//	10m - 10 * time.Minute
//	30s - 30 * time.Second
//
// So we can use readable string in our configuration files.
type TimeUnit struct {
	*valuex.NumUnit

	Duration time.Duration
}

// ParseTimeUnit
//
// Parse string value into TimeUnit Value. This function returns an error
// while the str value is not a valid time-unit string.
func ParseTimeUnit(str string) (*TimeUnit, error) {
	num, err := valuex.ParseNumUnit(str)
	if err != nil {
		return nil, err
	}

	if _, ok := gDurationMap[num.Unit]; !ok {
		return nil, fmt.Errorf("time unit '%v' not support", num.Unit)
	}

	duration := time.Duration(num.Value) * gDurationMap[num.Unit]
	return &TimeUnit{NumUnit: num, Duration: duration}, err
}
