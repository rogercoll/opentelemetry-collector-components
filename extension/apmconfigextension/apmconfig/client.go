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

type RemoteClient interface {
	// RemoteConfig returns the upstream remote configuration that needs to be applied. Empty RemoteConfig Attrs if no remote configuration is available for the specified service.
	RemoteConfig(context.Context, Params) (RemoteConfig, error)

	// LastConfig notifies the upstream server about the latest applied configuration.
	LastConfig(context.Context, Params, []byte) error
}

// RemoteConfig holds an agent remote configuration.
type RemoteConfig struct {
	Hash []byte

	Attrs map[string]string
}
