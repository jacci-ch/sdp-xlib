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

package valuex

// IntType - A type set of all integer type.
type IntType interface {
	int8 | int16 | int | int32 | int64
}

// UintType - A type set of all unsigned-integer type.
type UintType interface {
	uint8 | uint16 | uint | uint32 | uint64
}
