package tool

import (
	"context"
	"sync"

	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

type Factory[T ~string, R Provider] struct {
	namespaceMap map[T]*Namespace[T, R]
	lock         sync.RWMutex
	provider     R
}

type RunFunc func(ctx context.Context, args string) (*plugin.RunResponse, error)
type ContinueFunc func(ctx context.Context, session string) (*plugin.RunResponse, error)
type CacheGroupFunc func(args string) string
type ClientRendererFunc func(args string) string

type FactoryRegisterInfo[T ~string] struct {
	Name           T
	DisplayName    string
	Description    string
	Args           *schema.Args
	Run            RunFunc
	Continue       ContinueFunc
	CacheGroup     CacheGroupFunc
	ClientRenderer ClientRendererFunc
}

func (f *Factory[T, R]) RegisterNamespace(condition T, info *FactoryRegisterInfo[T]) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.namespaceMap[condition] = NewNamespace[T, R](f.provider, info)
}

func (f *Factory[T, R]) RegisterProvider(provider R) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.provider = provider
}

func (f *Factory[T, R]) Init() {
	for _, ns := range f.namespaceMap {
		Register(ns.Name(), func() Provider {
			return ns
		})
	}
}
