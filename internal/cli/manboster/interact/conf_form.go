package interact

import (
	"fmt"
	"sync"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/cli"
)

type handleFn func() error

var nilFunc handleFn = func() error { return nil }

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
	if fn == nil {
		return fmt.Errorf("no registered handler function for type %s", t)
	}
	return fn()
}

func handle[T ~string](p cli.Provider, f *configForm[T], options []cli.Option, title string, content string) error {
	var option cli.Option
	for {
		var err error
		option, err = p.Select(title, "Welcome to Manboster Configuration Wizard! Please choose which field you want to configure.", options, option.Value, func(option cli.Option) error {
			return nil
		})
		if err != nil {
			return err
		}

		err = f.Handle(T(option.Value))
		if err != nil {
			return err
		}

		if option.Value == "quit" {
			color.Yellow("Bye!")
			return nil
		}
	}
}
