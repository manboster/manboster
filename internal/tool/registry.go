package tool

import (
	"fmt"
	"regexp"
	"sync"
)

type ProviderFactory func() Provider

var (
	providerRegistry = make(map[string]ProviderFactory)
	mu               sync.RWMutex
)

var reg = regexp.MustCompile(`^[a-zA-Z0-9.-]+$`)

func Register(name string, factory ProviderFactory) {
	mu.Lock()
	defer mu.Unlock()
	// add error checker
	if !reg.MatchString(name) {
		panic(fmt.Sprintf("[Manboster Tool Registry] tool call value should not contain other characters except '-' and '.': %q", name))
	}
	if _, duplicant := providerRegistry[name]; duplicant {
		panic(fmt.Sprintf("[Manboster Tool Registry] Register called twice for provider %s", name))
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
