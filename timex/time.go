package timex

import (
	"time"
)

const (
	DateTime     = time.DateTime + ".000"
	DateTimeZone = time.DateTime + ".000+8:00"
)

type Time struct {
	time.Time
}

func Now() Time {
	return Time{Time: time.Now()}
}

func NowPtr() *Time {
	return &Time{Time: time.Now()}
}
