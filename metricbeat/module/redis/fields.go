// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package redis

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "redis", asset.ModuleFieldsPri, AssetRedis); err != nil {
		panic(err)
	}
}

// AssetRedis returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/redis.
func AssetRedis() string {
	return "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"
}
