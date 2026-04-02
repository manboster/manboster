package session

// NewManager creates an session manager instance.
func NewManager() *Manager {
	return &Manager{
		Sessions: make(map[string]Session),
	}
}

// GetSession gets session information for you.
func (m *Manager) GetSession(sessionId string) Session {
	m.Lock.RLock()
	defer m.Lock.RUnlock()

	session, avail := m.Sessions[sessionId]
	if !avail {
		return Session{}
	}
	return session
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
