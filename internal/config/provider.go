package config

import (
	"context"

	"github.com/charmbracelet/huh"
)

// Provider provides interfaces for all configurations
type Provider interface {
	Name() string
	DisplayName() string
	ToHuhGroup() []*huh.Group
	VerifyAndConvert(ctx context.Context) error
	GetConfig() any
}
