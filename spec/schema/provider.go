package schema

import (
	"context"
)

// PluginProvider defines what a tool will be, tool.Provider, plugin.Provider and skill.Provider is the implementation of this.
type PluginProvider interface {
	Name() string                                         // get the package name e.g. dev.manboster.websearch (We recommend you to create package names by reserved domain policy as this is the industrial standard)
	DisplayName() string                                  // get the display name of the provider
	MetaData() MetaData                                   // get the metadata of the plugin
	Requires() []RequirementType                          // get the requirement type of plugin
	Args() []Args                                         // get args description from the plugin
	Run(ctx context.Context, args string) (string, error) // passthrough by JSON
}
