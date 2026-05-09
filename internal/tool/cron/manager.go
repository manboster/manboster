package cron

import (
	"sync"
)

type Manager struct {
	lock   sync.RWMutex
	loaded map[string]LoadStatus
}

type LoadStatus struct {
	Loaded bool
	Entry  int
}

func NewManager() *Manager {
	return &Manager{
		loaded: make(map[string]LoadStatus),
		lock:   sync.RWMutex{},
	}
}

func (m *Manager) Loaded(name string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.loaded[name].Loaded
}

func (m *Manager) Load(name string, loaded bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	d, ok := m.loaded[name]
	if !ok {
		m.loaded[name] = LoadStatus{
			Loaded: loaded,
		}
		return
	}

	d.Loaded = loaded
	m.loaded[name] = d
}

func (m *Manager) GetEntry(name string) (int, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	data, ok := m.loaded[name]
	if !ok {
		return 0, false
	}
	return data.Entry, true
}

func (m *Manager) SetEntry(name string, entry int) {
	m.lock.Lock()
	defer m.lock.Unlock()

	data, ok := m.loaded[name]
	if !ok {
		m.loaded[name] = LoadStatus{
			Entry: entry,
		}
	}
	data.Entry = entry
	m.loaded[name] = data
}
