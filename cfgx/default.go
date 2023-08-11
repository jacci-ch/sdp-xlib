// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package cfgx

import (
	"github.com/jacci-ch/sdp-xlib/cfgx/cfgv"
)

// defaultValueKeeper
//
// The implementation of default-value interface.
type defaultValueKeeper struct {
}

func (k *defaultValueKeeper) GetValue(key string) (*cfgv.Value, bool) {
	return gValueKeeper.GetValue(Default, key)
}

func (k *defaultValueKeeper) ToInt64(key string, dst *int64, defaultValue int64) error {
	return gValueKeeper.ToInt64(Default, key, dst, defaultValue)
}

func (k *defaultValueKeeper) ToInt32(key string, dst *int32, defaultValue int32) error {
	return gValueKeeper.ToInt32(Default, key, dst, defaultValue)
}

func (k *defaultValueKeeper) ToInt(key string, dst *int, defaultValue int) error {
	return gValueKeeper.ToInt(Default, key, dst, defaultValue)
}

func (k *defaultValueKeeper) ToBool(key string, dst *bool, defaultValue bool) error {
	return gValueKeeper.ToBool(Default, key, dst, defaultValue)
}

func (k *defaultValueKeeper) ToStr(key string, dst *string, defaultValue string) error {
	return gValueKeeper.ToStr(Default, key, dst, defaultValue)
}

func (k *defaultValueKeeper) ToStrArray(key string, dst *[]string, defaultValue []string) error {
	return gValueKeeper.ToStrArray(Default, key, dst, defaultValue)
}
