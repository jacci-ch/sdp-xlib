package valuex

import (
	"sync/atomic"
	"unsafe"
)

func SetPointer(dst, src unsafe.Pointer) {
	atomic.StorePointer((*unsafe.Pointer)(dst), src)
}
