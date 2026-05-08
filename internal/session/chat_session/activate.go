package chat_session

import (
	"context"
)

func (m *Manager) Activate(sid string, cf context.CancelFunc) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	s, avail := m.Sessions[sid]
	if !avail {
		m.Sessions[sid] = Session{
			Active: true,
			Cancel: cf,
		}
	} else {
		s.Active = true
		s.Cancel = cf
		m.Sessions[sid] = s
	}
}

func (m *Manager) Deactivate(sid string) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	s, avail := m.Sessions[sid]
	if avail {
		s.Active = false
		s.Cancel = nil
	}
}
