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

import (
	"fmt"
	"github.com/jacci-ch/sdp-xlib/osx"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	gConfigFiles = []string{
		"./sdp.conf",
		"./conf/sdp.conf",
		"./conf/sdp/sdp.conf",
		"./cfg/sdp.conf",
		"./cfg/sdp/sdp.conf",
		"./etc/sdp.conf",
		"./etc/sdp/sdp.conf",
		"/etc/sdp.conf",
		"/etc/sdp/sdp.conf",
		"/user/local/etc/sdp.conf",
		"/user/local/etc/sdp/sdp.conf",
	}
)

// probeConfigFile - Detects and returns the existed file in the pre-set
// file list gConfigFiles. If the given file is not empty and not found
// this function will raise a panic.
func probeConfigFile(name string) (string, bool) {
	if len(name) != 0 {
		if osx.Exists(name) {
			return name, true
		} else if Cfg.PanicWithError {
			panic(fmt.Sprintf("cfgx: configuration file %v not found", name))
		}
	} else {
		for _, v := range gConfigFiles {
			if osx.Exists(v) {
				return v, true
			}
		}
	}

	return "", false
}

// loadConfigs - Load configuration items from a given file or probe the file
// in the pre-set list. We can set value to Cfg.CfgFile to tail this function
// to read from the user-defined configuration file.
func loadConfigs() (map[string]string, bool) {
	values := make(map[string]string)

	if name, ok := probeConfigFile(Cfg.CfgFile); ok {
		Cfg.CfgFile = name

		if iniFile, err := ini.Load(Cfg.CfgFile); err == nil {
			for _, section := range iniFile.Sections() {
				for _, key := range section.Keys() {
					values[key.Name()] = strings.TrimSpace(key.Value())
				}
			}

			return values, true
		} else if Cfg.PanicWithError {
			panic(fmt.Sprintf("cfgx: %v", err))
		}
	}

	return values, false
}
