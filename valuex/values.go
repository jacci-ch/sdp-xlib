// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package valuex

// StrPtr
//
// Returns the address of given string value.
func StrPtr(v string) *string {
	return &v
}

// Int64Ptr
//
// Returns the address of given int64 value.
func Int64Ptr(v int64) *int64 {
	return &v
}

// Int32Ptr
//
// Returns the address of given int32 value.
func Int32Ptr(v int32) *int32 {
	return &v
}

// IntPtr
//
// Returns the address of given int value.
func IntPtr(v int) *int {
	return &v
}

// BoolPtr
//
// Returns the address of given bool value.
func BoolPtr(v bool) *bool {
	return &v
}

// NotZeroIntPtr
//
// Detects the given pointer is not a non-zero pointer.
func NotZeroIntPtr(v *int) bool {
	return v != nil && *v != 0
}

// NotZeroInt32Ptr
//
// Detects the given pointer is not a non-zero pointer.
func NotZeroInt32Ptr(v *int32) bool {
	return v != nil && *v != 0
}

// NotZeroInt64Ptr
//
// Detects the given pointer is not a non-zero pointer.
func NotZeroInt64Ptr(v *int64) bool {
	return v != nil && *v != 0
}
