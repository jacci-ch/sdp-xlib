// Copyright 2023 to now() The SDP Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package multicast - implementation of multicast address resolver.
// We use
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
