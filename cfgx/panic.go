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

package cfgx

import "fmt"

// PanicIf - Raise a panic or returns the given value. We can
// set values to Cfg.PanicWithError to enable/disable panic.
func PanicIf[T any](ret T, src string, err error) (T, error) {
	if err != nil && Cfg.PanicWithError {
		panic(fmt.Sprintf("cfgx: invalid configuration value '%v', err = %v", src, err))
	}

	return ret, err
}
