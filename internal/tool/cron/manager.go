package cron

import "sync"

type Manager struct {
	lock   sync.RWMutex
	loaded map[string]bool
}

func NewManager() *Manager {
	return &Manager{
		loaded: make(map[string]bool),
		lock:   sync.RWMutex{},
	}
}

func (m *Manager) Loaded(name string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.loaded[name]
}

func (m *Manager) Load(name string, loaded bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.loaded[name] = loaded
}
