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

package apmconfig

import (
	"context"
)

// Params holds parameters for watching and notifying for config changes.
type Params struct {
	AgentUiD string

	// Service holds the name and optionally environment name used
	// for filtering the config to watch.
	Service struct {
		Name        string
		Environment string
	}
}

type Client interface {
	// RemoteConfig returns the upstream remote configuration that needs to be applied. Empty RemoteConfig Attrs if no remote configuration is available for the specified service.
	RemoteConfig(context.Context, Params) (RemoteConfig, error)

	// // LastConfig notifies the upstream server about the latest applied configuration.
	// LastConfig(context.Context, Params, []byte) error

	// Close the client's connection
	Close() error
}

// RemoteConfig holds an agent remote configuration.
type RemoteConfig struct {
	Hash []byte

	Attrs map[string]string
}
