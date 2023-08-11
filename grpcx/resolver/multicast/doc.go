// Copyright 2023 - now The SDP Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0 style
// license that can be found in the LICENSE file.

// Package multicast
//
// Implementation of multicast address resolver. We use
//
//	grpc.Dial("tcp://127.0.0.1:1234", opts...)
//
// directly to connect to a single target. If we want to
// connect to more than one target we can use this package
// and write code like:
//
//	grpc.Dial("multicast://127.0.0.1:1234,127.0.0.1:1235")
//
// all use the Dial function in this package like:
//
//	multicast.Dail([]string{"127.0.0.1:1234","127.0.0.1:1235"}, opts...)
//
// Caution: we should import multicast package before we prepare
// to use the multicast resolver.
package multicast
