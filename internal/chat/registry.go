package chat

import (
	"fmt"
	"sync"
)

type ProviderFactory func() Provider

var (
	providerRegistry = make(map[string]ProviderFactory)
	mu               sync.RWMutex
)

func Register(name string, factory ProviderFactory) {
	mu.Lock()
	defer mu.Unlock()
	if _, duplicant := providerRegistry[name]; duplicant {
		panic(fmt.Sprintf("chat: Register called twice for provider %s", name))
	}
	providerRegistry[name] = factory
}

func GetProvider(name string) (Provider, error) {
	mu.RLock()
	defer mu.RUnlock()
	factory, ok := providerRegistry[name]
	if !ok {
		return nil, fmt.Errorf("chat: unknown provider %q (did you forget to import it?)", name)
	}
	return factory(), nil
}

func AvailableProviders() []string {
	mu.RLock()
	defer mu.RUnlock()
	var list []string
	for name := range providerRegistry {
		list = append(list, name)
	}
	return list
}
