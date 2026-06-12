package search

import (
	"context"
	"fmt"
	"sync"
)

type searchHandleFunc func(context.Context, string) error

var (
	handlerMap  = make(map[string]searchHandleFunc)
	handlerLock sync.RWMutex
)

func RegisterProvider(name string, handler searchHandleFunc) {
	handlerLock.Lock()
	defer handlerLock.Unlock()
	if _, ok := handlerMap[name]; ok {
		panic("duplicate handler " + name)
	}
	handlerMap[name] = handler
}

func ExecProvider(ctx context.Context, name string, content string) error {
	handlerLock.RLock()
	defer handlerLock.RUnlock()
	if name == "auto" {
		for _, handler := range handlerMap {
			if err := handler(ctx, content); err != nil {
				return err
			}
			return nil
		}
	}
	handler, ok := handlerMap[name]
	if !ok {
		return fmt.Errorf("provider not found: " + name)
	}
	return handler(ctx, content)
}

func ListProviders() []string {
	handlerLock.RLock()
	defer handlerLock.RUnlock()
	var list []string
	for name := range handlerMap {
		list = append(list, name)
	}
	return list
}
