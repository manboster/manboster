package session

import "sync"

func (m *Manager) GetSessionLocks(sid string) *sync.Mutex {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	s, avail := m.SessionLocks[sid]
	if !avail || s == nil {
		m.SessionLocks[sid] = &sync.Mutex{}
		return m.SessionLocks[sid]
	}

	return s
}

func (m *Manager) DeleteSessionLocks(sid string) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	delete(m.SessionLocks, sid)
}
