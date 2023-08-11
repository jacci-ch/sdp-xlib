// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package valuex

import (
	"sync/atomic"
	"unsafe"
)

// SetPointer
//
// Assign a pointer value atomically to avoid the lock.
func SetPointer(dst, src unsafe.Pointer) {
	atomic.StorePointer((*unsafe.Pointer)(dst), src)
}
