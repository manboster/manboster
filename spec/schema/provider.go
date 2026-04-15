package schema

import (
	"context"
)

// PluginProvider defines what a tool will be, tool.Provider, plugin.Provider and skill.Provider is the implementation of this.
type PluginProvider interface {
	Name() string
	DisplayName() string
	MetaData() MetaData
	Requires() []RequirementType
	Args() []Args
	Run(ctx context.Context, args string) (string, error) // passthrough by JSON
}
