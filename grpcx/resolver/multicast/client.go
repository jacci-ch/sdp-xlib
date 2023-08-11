// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

package multicast

import (
	"fmt"
	"google.golang.org/grpc"
	"strings"
	"sync"
)

var (
	gLock sync.Mutex
	gConn *grpc.ClientConn
)

// Dial
//
// Connect to given targets listed in arguments endpoints.
// See doc.go for more information.
func Dial(endpoints []string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	target := fmt.Sprintf("%v://%v", Scheme, strings.Join(endpoints, Separator))
	return grpc.Dial(target, opts...)
}

// GetConn
//
// Returns the global grpc.ClientConn instance.
func GetConn(endpoints []string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	if gConn == nil {
		gLock.Lock()
		cc, err := Dial(endpoints, opts...)
		if err == nil {
			gConn = cc
		}
		gLock.Unlock()

		return cc, err
	}

	return gConn, nil
}
