package hachimi

import (
	"context"

	"github.com/manboster/manboster/internal/config"
)

// Provider defines hachimi's provider, only provide one model and easy to use. It is like llm's provider, it's much easier than llm one.
type Provider interface {
	Chat(ctx context.Context, sysMsg string, evaMsg string) (string, error)
	Init(ctx context.Context, config any) error
	Start(ctx context.Context) error
	Stop() error
	Name() string
	DisplayName() string
	Config() config.Provider
	New() Provider
}
