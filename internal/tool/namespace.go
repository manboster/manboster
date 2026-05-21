package tool

import (
	"context"

	"github.com/manboster/manboster/internal/engine/hook"
	"github.com/manboster/manboster/spec/config"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

type Namespace[T ~string, R Provider] struct {
	meta    schema.MetaData
	r       R
	regInfo *FactoryRegisterInfo[T]
}

func (n *Namespace[T, R]) Name() string {
	if n.regInfo == nil || n.regInfo.Name == "" {
		return n.meta.Name
	}
	return n.meta.Name + "." + string(n.regInfo.Name)
}

func (n *Namespace[T, R]) DisplayName() string {
	return n.meta.DisplayName
}

func (n *Namespace[T, R]) MetaData() schema.MetaData {
	return n.meta
}

func (n *Namespace[T, R]) Description() string {
	if n.regInfo.Description == "" {
		return n.meta.Description
	}
	return n.regInfo.Description
}

func (n *Namespace[T, R]) Requires() []schema.RequirementData {
	return n.meta.Requires
}

func (n *Namespace[T, R]) Args() *schema.Args {
	if n.regInfo == nil || n.regInfo.Args == nil {
		return n.r.Args()
	}
	return n.regInfo.Args
}

func (n *Namespace[T, R]) Init(ctx context.Context, conf any) error {
	return n.r.Init(ctx, conf)
}

func (n *Namespace[T, R]) Migrate(ctx context.Context, from int, conf any) (any, error) {
	return n.r.Migrate(ctx, from, conf)
}

func (n *Namespace[T, R]) CacheGroup(args string) string {
	if n.regInfo == nil || n.regInfo.CacheGroup == nil {
		return n.r.CacheGroup(args)
	}
	return n.regInfo.CacheGroup(args)
}

func (n *Namespace[T, R]) ClientRenderer(args string) string {
	if n.regInfo == nil || n.regInfo.ClientRenderer == nil {
		return n.r.ClientRenderer(args)
	}
	return n.regInfo.ClientRenderer(args)
}

func (n *Namespace[T, R]) Start(ctx context.Context) error {
	return n.r.Stop()
}

func (n *Namespace[T, R]) Run(ctx context.Context, args string) (*plugin.RunResponse, error) {
	if n.regInfo == nil || n.regInfo.Run == nil {
		return n.r.Run(ctx, args)
	}
	return n.regInfo.Run(ctx, args)
}

func (n *Namespace[T, R]) Continue(ctx context.Context, session string) (*plugin.RunResponse, error) {
	if n.regInfo == nil || n.regInfo.Continue == nil {
		return n.r.Continue(ctx, session)
	}
	return n.regInfo.Continue(ctx, session)
}

func (n *Namespace[T, R]) Config() config.Provider {
	return n.r.Config()
}

func (n *Namespace[T, R]) Stop() error {
	return n.r.Stop()
}

func (n *Namespace[T, R]) RegisterHook(registry *hook.Registry) {
	n.r.RegisterHook(registry)
}

func NewNamespace[T ~string, R Provider](r R, regInfo *FactoryRegisterInfo[T]) *Namespace[T, R] {
	return &Namespace[T, R]{
		r:       r,
		meta:    r.MetaData(),
		regInfo: regInfo,
	}
}
