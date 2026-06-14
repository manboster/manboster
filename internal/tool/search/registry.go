package search

import (
	"context"
	"fmt"
	"sync"
)

// Provider defines a search backend that can execute a query.
type Provider interface {
	Name() string
	DisplayName() string
	Search(ctx context.Context, query string) (string, error)
}

// Factory creates a new Provider instance.
type Factory func() Provider

var (
	factories     = make(map[string]Factory)
	active        = make(map[string]Provider)
	factoriesLock sync.RWMutex
	activeLock    sync.RWMutex
)

// Register adds a search provider factory to the registry.
func Register(name string, factory Factory) {
	factoriesLock.Lock()
	defer factoriesLock.Unlock()
	if _, ok := factories[name]; ok {
		panic(fmt.Sprintf("search: provider %q registered twice", name))
	}
	factories[name] = factory
}

// GetProvider returns a fresh provider instance from the factory registry.
func GetProvider(name string) (Provider, error) {
	factoriesLock.RLock()
	defer factoriesLock.RUnlock()
	factory, ok := factories[name]
	if !ok {
		return nil, fmt.Errorf("search: unknown provider %q", name)
	}
	return factory(), nil
}

// AvailProviders returns the names of all registered provider factories.
func AvailProviders() []string {
	factoriesLock.RLock()
	defer factoriesLock.RUnlock()
	list := make([]string, 0, len(factories))
	for name := range factories {
		list = append(list, name)
	}
	return list
}

// Activate marks a provider instance as ready for use.
func Activate(p Provider) {
	activeLock.Lock()
	defer activeLock.Unlock()
	active[p.Name()] = p
}

// Deactivate removes a provider instance from the active set.
func Deactivate(name string) {
	activeLock.Lock()
	defer activeLock.Unlock()
	delete(active, name)
}

// GetActive returns an active provider instance by name.
func GetActive(name string) (Provider, error) {
	activeLock.RLock()
	defer activeLock.RUnlock()
	p, ok := active[name]
	if !ok {
		return nil, fmt.Errorf("search: provider %q is not active", name)
	}
	return p, nil
}

// ActiveProviders returns the names of all currently active providers.
func ActiveProviders() []string {
	activeLock.RLock()
	defer activeLock.RUnlock()
	list := make([]string, 0, len(active))
	for name := range active {
		list = append(list, name)
	}
	return list
}
