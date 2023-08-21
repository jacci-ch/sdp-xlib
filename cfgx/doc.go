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

// Package cfgx
//
// This package provides functions to load configuration items in
// specified configuration file. The configuration file will be probed
// in the order of:
//
//	./sdp.conf
//	./conf/sdp.conf
//	./conf/sdp/sdp.conf
//	./cfg/sdp.conf
//	./cfg/sdp/sdp.conf
//	./etc/sdp.conf
//	./etc/sdp/sdp.conf
//	/etc/sdp.conf
//	/etc/sdp/sdp.conf
//	/user/local/etc/sdp.conf
//	/user/local/etc/sdp/sdp.conf
//
// When one of the file in the above list is found, the rest of other files
// will be ignored.
//
// We can use cfgx.AsInt() to parse configuration value as
// an int value. See cfgx.AsXxxx() for more information.
package cfgx
