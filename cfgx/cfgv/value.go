package cfgv

import (
	"errors"
	"strconv"
	"strings"
)

const (
	Default = "DEFAULT"
)

var (
	Empty = errors.New("empty")
)

type Value string

func (v Value) Addr() *Value {
	return &v
}

func (v Value) ToStr(dst *string) error {
	if len(v) == 0 {
		return Empty
	}

	*dst = string(v)
	return nil
}

func (v Value) ToStrArray(dst *[]string) error {
	for _, src := range strings.Split(string(v), ",") {
		if src = strings.TrimSpace(src); len(src) != 0 {
			*dst = append(*dst, src)
		}
	}

	if len(*dst) == 0 {
		return Empty
	}

	return nil
}

func (v Value) ToInt64(dst *int64) error {
	if len(v) == 0 {
		return Empty
	}

	value, err := strconv.ParseInt(string(v), 10, 64)
	if err != nil {
		return err
	}

	*dst = value
	return nil
}

func (v Value) ToInt32(dst *int32) error {
	var value int64
	if err := v.ToInt64(&value); err != nil {
		return err
	}

	*dst = int32(value)
	return nil
}

func (v Value) ToInt(dst *int) error {
	var value int64
	if err := v.ToInt64(&value); err != nil {
		return err
	}

	*dst = int(value)
	return nil
}

func (v Value) ToBool(dst *bool) error {
	if len(v) == 0 {
		return Empty
	}

	value, err := strconv.ParseBool(string(v))
	if err != nil {
		return err
	}

	*dst = value
	return nil
}
