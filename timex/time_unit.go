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

type TimeUnit struct {
	*valuex.NumUnit

	Duration time.Duration
}

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
