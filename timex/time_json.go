// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package timex

import (
	"fmt"
	"time"
)

// UnmarshalJSON
//
// Unmarshal Time from []byte value using default format string.
func (t *Time) UnmarshalJSON(bytes []byte) error {
	v, err := time.Parse(DateTimeZone, string(bytes))
	if err != nil {
		return err
	}

	t.Time = v
	return nil
}

// MarshalJSON
//
// Unmarshal Time to []byte value using default format string.
func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Format(DateTimeZone))), nil
}
