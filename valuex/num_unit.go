package valuex

import (
	"strconv"
	"strings"
)

type NumUnit struct {
	Value int
	Unit  string
}

func ParseNumUnit(str string) (*NumUnit, error) {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return nil, ErrEmpty
	}

	index := FirstNonDigit(str)
	v, err := strconv.ParseInt(str[0:index], 10, 64)
	if err != nil {
		return nil, err
	}

	return &NumUnit{Value: int(v), Unit: strings.TrimSpace(str[index:])}, nil
}
