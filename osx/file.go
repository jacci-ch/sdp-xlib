// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package osx

import "os"

// Exist
//
// Detect whether the file is existed or not. This function
// returns false if any errors.
func Exist(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}