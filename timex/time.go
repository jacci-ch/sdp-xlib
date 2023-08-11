// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package timex

import (
	"time"
)

const (
	DateTime     = time.DateTime + ".000"
	DateTimeZone = time.DateTime + ".000+8:00"
)

// Time
//
// A wrapper of time.Time defined in standard library.
type Time struct {
	time.Time
}

// Now
//
// Returns a Time object contains current time.
func Now() Time {
	return Time{Time: time.Now()}
}

// NowPtr
//
// Returns a pointer of Time object address contains current time.
func NowPtr() *Time {
	return &Time{Time: time.Now()}
}
