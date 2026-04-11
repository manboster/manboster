package session

import (
	"github.com/manboster/manboster/internal/llm"
)

func (m *Manager) AppendMessage(sid string, msg llm.Message) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	s, avail := m.Sessions[sid]
	if !avail {
		m.Sessions[sid] = Session{
			Messages: []llm.Message{msg},
		}
	} else {
		s.Messages = append(s.Messages, msg)
		m.Sessions[sid] = s
	}
}
