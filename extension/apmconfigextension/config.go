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

package apmconfigextension

import (
	"go.opentelemetry.io/collector/component"
)

type Config struct {
	CentralConfig CentralConfig `mapstructure:"centralconfig"`
	OpAMP         OpAMPConfig   `mapstructure:"opamp"`
}

type CentralConfig struct {
	Elastic ElasticConfig `mapstructure:"elastic"`
}

type ElasticConfig struct {
	Apm struct {
		Server struct {
			URLs []string `mapstructure:"urls"`
			// TODO add timeout
		} `mapstructure:"server"`
		SecretToken string `mapstructure:"secret_token"`
	} `mapstructure:"apm"`
}

type OpAMPConfig struct {
	Server OpAMPServerConfig `mapstructure:"server"`
}

type OpAMPServerConfig struct {
	Endpoint string `mapstructure:"endpoint"`
}

var _ component.Config = (*Config)(nil)
