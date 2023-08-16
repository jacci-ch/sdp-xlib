// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package redisx

import (
	"github.com/jacci-ch/sdp-xlib/logx"
	"sync/atomic"
	"unsafe"
)

// init
//
// Load configurations and create init the global values.
func init() {
	cfg, err := LoadConfigs()
	if err != nil {
		logx.Logger.Fatal(err)
		return
	}

	if err = cfg.Validate(); err != nil {
		logx.Logger.Fatal(err)
		return
	}

	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&Cfg)), unsafe.Pointer(cfg))
	GetClient() // Create the default redis client
}
