package telegram

import "sync"

type Manager struct {
	avail bool
	lock  sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		avail: false,
		lock:  sync.RWMutex{},
	}
}

func (m *Manager) Avail() bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.avail
}

func (m *Manager) SetAvail(avail bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.avail = avail
}
