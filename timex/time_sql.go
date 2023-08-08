package timex

import (
	"database/sql/driver"
	"errors"
	"time"
)

func (t *Time) Value() (value driver.Value, err error) {
	if t == nil || t.UnixNano() == (time.Time{}).UnixNano() {
		return nil, nil
	}

	return t.Format(time.DateTime), nil
}

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
