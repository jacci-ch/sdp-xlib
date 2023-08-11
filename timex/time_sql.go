// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package timex

import (
	"database/sql/driver"
	"errors"
	"time"
)

// Value
//
// Encode Time value to sql driver.Value for database operations.
func (t *Time) Value() (value driver.Value, err error) {
	if t == nil || t.UnixNano() == (time.Time{}).UnixNano() {
		return nil, nil
	}

	return t.Format(time.DateTime), nil
}

// Scan
//
// Decode Time value from sql driver.Value in database operations.
func (t *Time) Scan(i any) error {
	switch i.(type) {
	case []byte:
		if v, err := time.Parse(time.DateTime, string(i.([]byte))); err == nil {
			t.Time = v
		} else {
			return errors.New("timex: " + err.Error())
		}
	}

	return nil
}
