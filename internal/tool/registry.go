package tool

import (
	"fmt"
	"sync"
)

var IsLoading = false

type ProviderFactory func() Provider

var (
	providerRegistry = make(map[string]ProviderFactory)
	importedRegistry = make(map[string]bool)
	mu               sync.RWMutex
)

func Register(name string, factory ProviderFactory) {
	mu.Lock()
	defer mu.Unlock()
	if _, duplicant := providerRegistry[name]; duplicant {
		panic(fmt.Sprintf("[Manboster Chat Registry] Register called twice for provider %s", name))
	}
	providerRegistry[name] = factory
}

func GetProvider(name string) (Provider, error) {
	mu.RLock()
	defer mu.RUnlock()
	factory, ok := providerRegistry[name]
	if !ok {
		return nil, fmt.Errorf("tool: unknown provider %q (did you forget to import it?)", name)
	}
	if _, valid := importedRegistry[name]; valid {
		return nil, fmt.Errorf("tool: you define 2 times with %q! The second one will be ignored", name)
	}
	return factory(), nil
}

// AvailProviders gets providers to iterate
func AvailProviders() []string {
	mu.RLock()
	defer mu.RUnlock()
	var list []string
	for name := range providerRegistry {
		list = append(list, name)
	}
	return list
}
