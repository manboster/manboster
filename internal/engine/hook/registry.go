package hook

import "sync"

type Registry struct {
	providers map[string]map[string]any
	lock      sync.RWMutex
}

func NewRegistry() *Registry {
	return &Registry{
		providers: make(map[string]map[string]any),
		lock:      sync.RWMutex{},
	}
}

func (r *Registry) Register(t EngineHookType, name string, provider any) {
	r.lock.Lock()
	defer r.lock.Unlock()

	ma, avail := r.providers[string(t)]
	if !avail || ma == nil {
		r.providers[string(t)] = map[string]any{
			name: provider,
		}
		return
	}

	if ma[name] != nil {
		return
	}

	ma[name] = provider
	r.providers[string(t)] = ma
}

func (r *Registry) GetProviders(t EngineHookType) []any {
	r.lock.RLock()
	defer r.lock.RUnlock()
	m, avail := r.providers[string(t)]
	if !avail {
		return nil
	}

	var ret []any
	for _, ma := range m {
		ret = append(ret, ma)
	}
	return ret
}

var Reg = NewRegistry()
