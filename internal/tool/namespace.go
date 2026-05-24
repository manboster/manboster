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
	return n.meta.Name
}

func (n *Namespace[T, R]) DisplayName() string {
	return n.meta.DisplayName
}

func (n *Namespace[T, R]) MetaData() schema.MetaData {
	return n.meta
}

func (n *Namespace[T, R]) Description() string {
	return n.meta.Description
}

func (n *Namespace[T, R]) Requires() []schema.RequirementData {
	return n.meta.Requires
}

func (n *Namespace[T, R]) Args() *schema.Args {
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
		meta:    CreateNewMetaData[T](r.MetaData(), regInfo),
		regInfo: regInfo,
	}
}

func CreateNewMetaData[T ~string](meta schema.MetaData, regInfo *FactoryRegisterInfo[T]) schema.MetaData {
	forkedMeta := meta
	if regInfo != nil {
		if regInfo.Meta.Name != "" {
			forkedMeta.Name += "." + regInfo.Meta.Name
		}
		if regInfo.Meta.DisplayName != "" {
			forkedMeta.DisplayName = regInfo.Meta.DisplayName
		}
		if regInfo.Meta.Description != "" {
			forkedMeta.Description = regInfo.Meta.Description
		}
		if regInfo.Meta.MinUserType != "" {
			forkedMeta.MinUserType = regInfo.Meta.MinUserType
		}
		if regInfo.Meta.Represent != "" {
			forkedMeta.Represent = regInfo.Meta.Represent
		}
		forkedMeta.Irreversible = regInfo.Meta.Irreversible
	}
	return forkedMeta
}
