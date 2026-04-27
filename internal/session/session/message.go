package session

import (
	"github.com/manboster/manboster/spec/llm"
)

// GetMessages return messages from sid
func (m *Manager) GetMessages(sid string) []llm.Message {
	m.Lock.RLock()
	defer m.Lock.RUnlock()

	s, avail := m.Sessions[sid]
	if !avail {
		return nil
	}

	var msgs []llm.Message
	for _, event := range s.Events {
		if event.EventType&llm.EventMessage != 0 && event.Message != nil {
			msgs = append(msgs, *event.Message)
		}
	}
	return msgs
}
