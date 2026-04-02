package config

import (
	"fmt"
	"strings"
	"sync"
)

type ProviderFactory func() Provider

var (
	providerRegistry = make(map[string]ProviderFactory)
	mu               sync.RWMutex
)

// Register helps instance register their services
func Register(name string, factory ProviderFactory) {
	mu.Lock()
	defer mu.Unlock()
	if _, duplicant := providerRegistry[name]; duplicant {
		panic(fmt.Sprintf("Config register called twice for provider %s", name))
	}
	providerRegistry[name] = factory
}

// GetProvider gets providers to users
func GetProvider(name string) (Provider, error) {
	mu.RLock()
	defer mu.RUnlock()
	factory, ok := providerRegistry[name]
	if !ok {
		return nil, fmt.Errorf("unknown config provider %q (did you forget to import it?)", name)
	}
	return factory(), nil
}

// AvailProviders gets providers to iterate
func AvailProviders(scope string) []string {
	mu.RLock()
	defer mu.RUnlock()
	var list []string
	for name := range providerRegistry {
		if strings.HasPrefix(name, scope) && strings.Split(name, ":")[0] == scope {
			list = append(list, name)
		}
	}
	return list
}
