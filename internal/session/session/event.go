package session

import (
	"github.com/manboster/manboster/spec/llm"
)

func (m *Manager) AppendEvent(sid string, event llm.Event) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	s, avail := m.Sessions[sid]
	if !avail {
		m.Sessions[sid] = Session{
			Events: []llm.Event{
				event,
			},
		}
	} else {
		s.Events = append(s.Events, event)
		m.Sessions[sid] = s
	}
}
