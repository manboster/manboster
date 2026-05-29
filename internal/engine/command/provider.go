package command

import (
	"context"
	"sync"

	"github.com/manboster/manboster/spec/chat"
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
	_, ok := p.funcMap[t]
	if ok {
		panic("[Manboster Command Provider] provider already registered!!!")
	}
	p.funcMap[t] = fn
}

func (p *Provider[T]) Handle(ctx context.Context, t T, instance chat.Provider, msg *chat.Message, sessionId string) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	fn, ok := p.funcMap[t]
	if !ok {
		return p.defaultFunc(ctx, instance, msg, sessionId)
	}

	return fn(ctx, instance, msg, sessionId)
}

func (p *Provider[T]) Default(fn handleFunc) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.defaultFunc = fn
}
