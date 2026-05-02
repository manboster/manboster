package chat_session

import "context"

func (m *Manager) SetSessionCancel(sid string, cancel context.CancelFunc) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	s, avail := m.Sessions[sid]
	if !avail {
		m.Sessions[sid] = Session{
			SessCancel: cancel,
		}
		return
	}
	s.SessCancel = cancel
	m.Sessions[sid] = s
	return
}

func (m *Manager) SessionCancel(sid string) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	s, avail := m.Sessions[sid]
	if !avail {
		return
	}
	if s.SessCancel != nil {
		s.SessCancel()
	}
}
