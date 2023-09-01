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

import (
	"reflect"

	"github.com/jacci-ch/sdp-xlib/logx"
)

func SetField(dst any, name string, v2 any) {
	rVal := reflect.Indirect(reflect.ValueOf(dst))
	rField := rVal.FieldByName(name)

	if rField.IsValid() && rField.Type().String() == reflect.TypeOf(v2).String() {
		if rField.CanSet() {
			rField.Set(reflect.ValueOf(v2))
		} else {
			logx.Warnf("field %v can't be set", name)
		}
	} else {
		logx.Warnf("filed %v is invalid or type miss-matched")
	}
}
