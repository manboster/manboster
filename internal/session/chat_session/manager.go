package chat_session

import (
	"fmt"
	"sync"
)

// New creates a session manager instance.
func New() *Manager {
	return &Manager{
		Sessions: make(map[string]Session),
		Lock:     sync.RWMutex{},
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

func (m *Manager) AvailSession(sessionId string) bool {
	m.Lock.RLock()
	defer m.Lock.RUnlock()
	_, avail := m.Sessions[sessionId]
	return avail
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

	s, avail := m.Sessions[key]
	if !avail {
		return
	}
	if s.Cancel != nil {
		s.Cancel()
	}
	if s.SessCancel != nil {
		s.SessCancel()
	}
	if s.Ch != nil {
		close(s.Ch)
	}
	delete(m.Sessions, key)
}

// ID helps you to get sessionID
func (m *Manager) ID(provider string, chatId string) string {
	return fmt.Sprintf("%s:%s", provider, chatId)
}
