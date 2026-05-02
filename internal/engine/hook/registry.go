package hook

import "sync"

type Registry struct {
	providers map[string][]any
	lock      sync.RWMutex
}

func NewRegistry() *Registry {
	return &Registry{
		providers: make(map[string][]any),
		lock:      sync.RWMutex{},
	}
}

func (r *Registry) Register(t EngineHookType, provider any) {
	r.lock.Lock()
	defer r.lock.Unlock()

	list, avail := r.providers[string(t)]
	if !avail {
		r.providers[string(t)] = []any{provider}
		return
	}
	list = append(list, provider)
	r.providers[string(t)] = list
}

func (r *Registry) GetProviders(t EngineHookType) []any {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.providers[string(t)]
}

var Reg = NewRegistry()
