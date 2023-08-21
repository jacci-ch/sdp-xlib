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

var (
	Cfg    *Config
	DefCfg = &Config{
		PanicWithError: true,
	}
)

type Config struct {
	// The user-specified configuration file name.
	// We can do parse os.args to pick the configuration file arguments.
	CfgFile string

	// Whether raise a panic when errors found or not. If PanicWithError is true
	// this module, cfgx, will raise a panic and exit the program when errors found.
	//
	// Otherwise, cfgx do nothing but returns a empty value set.
	PanicWithError bool
}
