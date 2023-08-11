// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package cfgx

var (
	Def = &defaultValueKeeper{}
)

// Export cfgv.ValueGetter methods functions as package functions.
// All functions calls the global value keeper methods.

func ToInt64(name, key string, dst *int64, defaultValue int64) error {
	return gValueKeeper.ToInt64(name, key, dst, defaultValue)
}

func ToInt32(name, key string, dst *int32, defaultValue int32) error {
	return gValueKeeper.ToInt32(name, key, dst, defaultValue)
}

func ToInt(name, key string, dst *int, defaultValue int) error {
	return gValueKeeper.ToInt(name, key, dst, defaultValue)
}

func ToBool(name, key string, dst *bool, defaultValue bool) error {
	return gValueKeeper.ToBool(name, key, dst, defaultValue)
}

func ToStr(name, key string, dst *string, defaultValue string) error {
	return gValueKeeper.ToStr(name, key, dst, defaultValue)
}

func ToStrArray(name, key string, dst *[]string, defaultValue []string) error {
	return gValueKeeper.ToStrArray(name, key, dst, defaultValue)
}
