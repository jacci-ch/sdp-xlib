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

package redisx

import "github.com/redis/go-redis/v9"

var (
	Client redis.UniversalClient
)

// NewClient - creates a new redis Universal Client. See redis documents for more
// about redis.UniversalClient.
func NewClient(cfg *Config) (redis.UniversalClient, error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      cfg.Addrs,
		Username:   cfg.Username,
		Password:   cfg.Password,
		ClientName: cfg.ClientName,
		MasterName: cfg.MasterName,
		DB:         cfg.Database,
	}), nil
}
