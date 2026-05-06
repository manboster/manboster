package plugin

import (
	"context"

	"github.com/manboster/manboster/spec/config"
	"github.com/manboster/manboster/spec/schema"
)

// Provider defines what a tool will be, tool.Provider, plugin.Provider and skill.Provider is the implementation of this.
type Provider interface {
	Name() string                                                       // get the package name e.g. dev.manboster.websearch (We recommend you to create package names by reserved domain policy as this is the industrial standard)
	DisplayName() string                                                // get the display name of the provider
	MetaData() schema.MetaData                                          // get the metadata of the plugin
	Requires() []schema.RequirementData                                 // get the requirement type of plugin
	Args() *schema.Args                                                 // get args description from the plugin
	Init(ctx context.Context, conf any) error                           // initialize the plugin
	Migrate(ctx context.Context, from int, conf any) (any, error)       // migration when the version is too low
	CacheGroup(args string) string                                      // return this action's group as a string, as this string will be used to group actions and control what should do next
	Start(ctx context.Context) error                                    // if long polling, it would work
	Run(ctx context.Context, args string) (*RunResponse, error)         // passthrough by JSON
	Continue(ctx context.Context, session string) (*RunResponse, error) // continue to do, avoid to interrupt
	Config() config.Provider                                            // if config.Provider is nil, it means there is no need to configure.
	Stop() error                                                        // force stop
}
