package llm

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
	providerRegistry[name] = factory
}

func GetProvider(name string) (Provider, error) {
	mu.RLock()
	defer mu.RUnlock()
	factory, ok := providerRegistry[name]
	if !ok {
		return nil, fmt.Errorf("llm: unknown provider %q (did you forget to import it?)", name)
	}
	return factory(), nil
}

func GetLLMProviders() []string {
	mu.RLock()
	defer mu.RUnlock()
	var list []string
	for name := range providerRegistry {
		list = append(list, name)
	}
	return list
}
