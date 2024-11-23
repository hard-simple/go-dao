package config

import "context"

type (

	// Config is configuration of DAO.
	// Each DB implementation provider decides how to implement it.
	Config interface {
	}

	// Configurable allows use set up configuration for an instance of definition.
	Configurable interface {
		Configure(ctx context.Context, config Config) error
	}

	// Producer produces config for a specific DAO implementation.
	Producer func(ctx context.Context) Config
)
