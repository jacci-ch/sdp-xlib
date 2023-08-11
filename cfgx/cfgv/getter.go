// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package cfgv

// ValueGetter
//
// An interface for configuration value getter. You can use the function
// ValueGetter.GetValue(name, key) to get a Value object. The arguments:
//
//	name - the Section name
//	key  - The key in configuration file.
//
// Or you can use function ValueGetter.ToInt(nme, key) (or other ToXxx() function)
// to fetch the value and convert it to an int (or other values) value.
type ValueGetter interface {
	GetValue(name, key string) (*Value, bool)

	ToInt64(name, key string, dst *int64, defVal int64) error
	ToInt32(name, key string, dst *int32, defVal int32) error
	ToInt(name, key string, dst *int, defVal int) error
	ToBool(name, key string, dst *bool, defVal bool) error
	ToStr(name, key string, dst *string, defVal string) error
	ToStrArray(name, key string, dst *[]string, defVal []string) error
}

// DefaultValueGetter
//
// A convenience interface for caller to use DEFAULT keys in default section.
// The section name is no more need to use functions in this interface cause
// the default section name is used instead.
type DefaultValueGetter interface {
	GetValue(key string) (*Value, bool)

	ToInt64(key string, dst *int64, defVal int64) error
	ToInt32(key string, dst *int32, defVal int32) error
	ToInt(key string, dst *int, defVal int) error
	ToBool(key string, dst *bool, defVal bool) error
	ToStr(key string, dst *string, defVal string) error
	ToStrArray(key string, dst *[]string, defVal []string) error
}
