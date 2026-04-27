package session

// SetModel sets model information of a session
func (m *Manager) SetModel(sid string, provider string, model string) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	s, avail := m.Sessions[sid]
	if !avail {
		return
	}
	s.Model = model
	s.Provider = provider
	m.Sessions[sid] = s
}

// GetModel gets model information of a session
func (m *Manager) GetModel(sid string) (string, string, bool) {
	m.Lock.RLock()
	defer m.Lock.RUnlock()

	s, avail := m.Sessions[sid]
	if !avail {
		return "", "", false
	}
	return s.Provider, s.Model, true
}
