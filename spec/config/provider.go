package config

import (
	"context"

	"github.com/manboster/manboster/spec/cli"
)

// Provider provides interfaces for all configurations
type Provider interface {
	Name() string
	DisplayName() string
	Args() *Args
	Validate() error
	GetConfig() any
}

type ProviderWithSetup interface {
	Provider
	Setup(ctx context.Context, p cli.Provider) error
}
