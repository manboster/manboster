package chat_session

import "github.com/manboster/manboster/spec/chat"

func (m *Manager) GetChan(sid string) chan *chat.Message {
	m.Lock.RLock()
	defer m.Lock.RUnlock()
	s, avail := m.Sessions[sid]
	if !avail {
		return nil
	}
	return s.Ch
}

func (m *Manager) CreateChan(sid string, ch chan *chat.Message) {
	m.Lock.Lock()
	defer m.Lock.Unlock()
	s, avail := m.Sessions[sid]
	if !avail {
		se := Session{
			Ch: ch,
		}
		m.Sessions[sid] = se
		return
	}
	s.Ch = ch
	m.Sessions[sid] = s
	return
}

func (m *Manager) AvailChan(sid string) bool {
	m.Lock.RLock()
	defer m.Lock.RUnlock()
	s, avail := m.Sessions[sid]
	if !avail {
		return false
	}
	return s.Ch != nil
}
