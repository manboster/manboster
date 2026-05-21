package command

import (
	"context"
	"sync"
)

type Provider[T ~string] struct {
	funcMap     map[T]handleFunc
	defaultFunc handleFunc
	lock        sync.RWMutex
}

func NewProvider[T ~string]() *Provider[T] {
	return &Provider[T]{
		funcMap: make(map[T]handleFunc),
		lock:    sync.RWMutex{},
	}
}

func (p *Provider[T]) Register(t T, fn handleFunc) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.funcMap[t] = fn
}

func (p *Provider[T]) Handle(ctx context.Context, t T) error {
	fn, ok := p.funcMap[t]
	if !ok {
		return p.defaultFunc(ctx)
	}

	return fn(ctx)
}

func (p *Provider[T]) Default(fn handleFunc) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.defaultFunc = fn
}
