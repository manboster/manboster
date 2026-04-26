package llm

import (
	"fmt"
	"sync"

	"github.com/manboster/manboster/spec/llm"
)

type ProviderFactory func() llm.Provider

var (
	providerRegistry = make(map[string]ProviderFactory)
	mu               sync.RWMutex
)

func Register(name string, factory ProviderFactory) {
	mu.Lock()
	defer mu.Unlock()
	providerRegistry[name] = factory
}

func GetProvider(name string) (llm.Provider, error) {
	mu.RLock()
	defer mu.RUnlock()
	factory, ok := providerRegistry[name]
	if !ok {
		return nil, fmt.Errorf("llm: unknown provider %q (did you forget to import it?)", name)
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

// AllProviders gets all providers back
func AllProviders() []llm.Provider {
	mu.RLock()
	defer mu.RUnlock()
	var list []llm.Provider
	for p := range providerRegistry {
		list = append(list, providerRegistry[p]())
	}
	return list
}
