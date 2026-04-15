package tool

import "context"

// Provider defines what a tool will be
type Provider interface {
	Name() string
	Description() string
	Args() map[string]any
	Execute(ctx context.Context, args string) (string, error) // passthrough by JSON
}
