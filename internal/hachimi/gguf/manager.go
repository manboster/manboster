package gguf

import "sync"

type Manager struct {
	avail      bool
	availModel bool
	lock       sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		avail:      false,
		availModel: false,
		lock:       sync.RWMutex{},
	}
}

func (m *Manager) Avail() bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.avail
}

func (m *Manager) AvailModel() bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.availModel
}

func (m *Manager) SetAvail(avail bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.avail = avail
}

func (m *Manager) SetAvailModel(avail bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.availModel = avail
}

func (m *Manager) IsReady() bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.avail && m.availModel
}
