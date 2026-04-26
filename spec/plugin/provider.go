package plugin

import (
	"context"

	"github.com/manboster/manboster/spec/schema"
)

// Provider defines what a tool will be, tool.Provider, plugin.Provider and skill.Provider is the implementation of this.
type Provider interface {
	Name() string                                                       // get the package name e.g. dev.manboster.websearch (We recommend you to create package names by reserved domain policy as this is the industrial standard)
	DisplayName() string                                                // get the display name of the provider
	MetaData() schema.MetaData                                          // get the metadata of the plugin
	Requires() []schema.RequirementData                                 // get the requirement type of plugin
	Args() *schema.Args                                                 // get args description from the plugin
	Init(ctx context.Context) error                                     // initialize the plugin
	Start(ctx context.Context) error                                    // if long polling, it would work
	Run(ctx context.Context, args string) (*RunResponse, error)         // passthrough by JSON
	Continue(ctx context.Context, session string) (*RunResponse, error) // continue to do, avoid to interrupt
	Close() error                                                       // force stop
}
