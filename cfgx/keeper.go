// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package cfgx

import (
	"github.com/jacci-ch/sdp-xlib/cfgx/cfgv"
)

const (
	Default = "DEFAULT"
)

var (
	gValueKeeper ValueKeeper
)

// ValueKeeper
//
// A map to store all configuration values. The tree: map[<section>][<key>]value.
// ValueKeeper also implements the cfgv.ValueGetter interface.
type ValueKeeper map[string]map[string]*cfgv.Value

// NewValueKeeper
//
// The ValueKeeper creator for convenience.
func NewValueKeeper() ValueKeeper {
	return make(ValueKeeper)
}

func (k ValueKeeper) GetValue(name, key string) (*cfgv.Value, bool) {
	if section, ok := gValueKeeper[name]; ok {
		if value, ok := section[key]; ok {
			return value, true
		}
	}

	if section, ok := gValueKeeper[Default]; ok {
		if value, ok := section[key]; ok {
			return value, true
		}
	}

	return nil, false
}

func (k ValueKeeper) ToInt64(name, key string, dst *int64, defaultValue int64) error {
	if value, ok := k.GetValue(name, key); !ok {
		*dst = defaultValue
		return nil
	} else if err := value.ToInt64(dst); err == nil {
		return nil
	} else if err == cfgv.Empty {
		*dst = defaultValue
		return nil
	} else {
		return err
	}
}

func (k ValueKeeper) ToInt32(name, key string, dst *int32, defaultValue int32) error {
	if value, ok := k.GetValue(name, key); !ok {
		*dst = defaultValue
		return nil
	} else if err := value.ToInt32(dst); err == nil {
		return nil
	} else if err == cfgv.Empty {
		*dst = defaultValue
		return nil
	} else {
		return err
	}
}

func (k ValueKeeper) ToInt(name, key string, dst *int, defaultValue int) error {
	if value, ok := k.GetValue(name, key); !ok {
		*dst = defaultValue
		return nil
	} else if err := value.ToInt(dst); err == nil {
		return nil
	} else if err == cfgv.Empty {
		*dst = defaultValue
		return nil
	} else {
		return err
	}
}

func (k ValueKeeper) ToBool(name, key string, dst *bool, defaultValue bool) error {
	if value, ok := k.GetValue(name, key); !ok {
		*dst = defaultValue
		return nil
	} else if err := value.ToBool(dst); err == nil {
		return nil
	} else if err == cfgv.Empty {
		*dst = defaultValue
		return nil
	} else {
		return err
	}
}

func (k ValueKeeper) ToStr(name, key string, dst *string, defaultValue string) error {
	if value, ok := k.GetValue(name, key); !ok {
		*dst = defaultValue
		return nil
	} else if err := value.ToStr(dst); err == nil {
		return nil
	} else if err == cfgv.Empty {
		*dst = defaultValue
		return nil
	} else {
		return err
	}
}

func (k ValueKeeper) ToStrArray(name, key string, dst *[]string, defaultValue []string) error {
	if value, ok := k.GetValue(name, key); !ok {
		*dst = defaultValue
		return nil
	} else if err := value.ToStrArray(dst); err == nil {
		return nil
	} else if err == cfgv.Empty {
		*dst = defaultValue
		return nil
	} else {
		return err
	}
}
