package chat

import (
	"sync"
)

func (m *Manager) GetSessionChatLocks(sid string) *sync.Mutex {
	s, avail := m.SessionChatLocks[sid]
	if !avail || s == nil {
		m.SessionChatLocks[sid] = &sync.Mutex{}
		return m.SessionChatLocks[sid]
	}

	return s
}

func (m *Manager) DeleteSessionChatLocks(sid string) {
	delete(m.SessionChatLocks, sid)
}
