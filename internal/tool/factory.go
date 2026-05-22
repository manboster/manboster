package tool

import (
	"sync"

	"github.com/manboster/manboster/spec/schema"
)

type Factory[T ~string, R Provider] struct {
	namespaceMap map[T]*Namespace[T, R]
	lock         sync.RWMutex
	provider     R
}

type FactoryRegisterInfo[T ~string] struct {
	Meta           schema.MetaData
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

func NewFactory[T ~string, R Provider]() *Factory[T, R] {
	return &Factory[T, R]{
		namespaceMap: make(map[T]*Namespace[T, R]),
		lock:         sync.RWMutex{},
	}
}
