package chat_session

import "github.com/manboster/manboster/spec/chat"

func (m *Manager) LoadOrCreateChan(sid string) (chan *chat.Message, bool) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	s, avail := m.Sessions[sid]
	if !avail {
		ch := make(chan *chat.Message, 16)
		se := Session{
			Ch: ch,
		}
		m.Sessions[sid] = se
		return ch, true
	}
	if s.Ch == nil {
		s.Ch = make(chan *chat.Message, 16)
		m.Sessions[sid] = s
		return s.Ch, true
	}
	return s.Ch, false
}
