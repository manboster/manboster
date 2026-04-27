package session

import "sync"

func (m *Manager) GetSessionChatLocks(sid string) *sync.Mutex {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	s, avail := m.SessionChatLocks[sid]
	if !avail || s == nil {
		m.SessionChatLocks[sid] = &sync.Mutex{}
		return m.SessionChatLocks[sid]
	}

	return s
}

func (m *Manager) DeleteSessionChatLocks(sid string) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	delete(m.SessionChatLocks, sid)
}
