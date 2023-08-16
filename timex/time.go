// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package timex

import (
	"github.com/jacci-ch/sdp-xlib/valuex"
	"time"
)

const (
	DateTime      = time.DateTime + ".000"
	DateTimeZone  = time.DateTime + ".000+8:00"
	DefaultFormat = DateTimeZone
)

// Time
//
// A wrapper of time.Time defined in standard library.
type Time struct {
	time.Time
}

// Fmt
//
// Format Time to string value with default layout string.
func (t *Time) Fmt() string {
	return t.Format(DefaultFormat)
}

// FmtPtr
//
// Format Time to string value with default layout string, and
// returns the pointer of the formatted string.
func (t *Time) FmtPtr() *string {
	return valuex.StrPtr(t.Fmt())
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

// Parse
//
// Decode string value to Time structure.
func Parse(str string) (Time, error) {
	val, err := time.Parse(DefaultFormat, str)
	return Time{Time: val}, err
}

// ParsePtr
//
// Similar with Parse but return a Time pointer.
func ParsePtr(str string) (*Time, error) {
	val, err := time.Parse(DefaultFormat, str)
	if err != nil {
		return nil, err
	}

	return &Time{Time: val}, nil
}
