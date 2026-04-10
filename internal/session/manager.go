package session

import "fmt"

// NewManager creates an session manager instance.
func NewManager() *Manager {
	return &Manager{
		Sessions: make(map[string]Session),
	}
}

// GetSession gets session information for you.
func (m *Manager) GetSession(sessionId string) (Session, bool) {
	m.Lock.RLock()
	defer m.Lock.RUnlock()

	session, avail := m.Sessions[sessionId]
	if !avail {
		return Session{}, false
	}
	return session, true
}

// SetSession sets session information for you.
func (m *Manager) SetSession(key string, session Session) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	m.Sessions[key] = session
}

// DeleteSession helps you delete session.
func (m *Manager) DeleteSession(key string) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	delete(m.Sessions, key)
}

// ID helps you to get sessionID
func (m *Manager) ID(provider string, chatId string) string {
	return fmt.Sprintf("%s:%s", provider, chatId)
}
