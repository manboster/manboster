package interact

import (
	"fmt"
	"sync"
)

type handleFn func() error

type ConfigForm[T ~string] interface {
	Register(t T, handlerFunc handleFn)
	Handle(t T) error
}

type configForm[T ~string] struct {
	lock   sync.RWMutex
	regMap map[T]handleFn
}

func newConfigForm[T ~string]() *configForm[T] {
	return &configForm[T]{
		lock:   sync.RWMutex{},
		regMap: make(map[T]handleFn),
	}
}

func (f *configForm[T]) Register(t T, handlerFunc handleFn) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.regMap[t] = handlerFunc
}

func (f *configForm[T]) Handle(t T) error {
	f.lock.RLock()
	defer f.lock.RUnlock()

	fn, ok := f.regMap[t]
	if !ok {
		return fmt.Errorf("no registered handler function for type %s", t)
	}
	return fn()
}
