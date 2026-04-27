package selection

import (
	"github.com/manboster/manboster/spec/chat"
)

func (m *Manager) GetSelectionChan(sid string) chan *chat.Message {
	m.SelectionLock.Lock()
	defer m.SelectionLock.Unlock()
	c, ok := m.SelectionChan[sid]
	if !ok {
		return nil
	}
	return c
}

func (m *Manager) SetSelectionChan(sid string, c chan *chat.Message) {
	m.SelectionLock.Lock()
	defer m.SelectionLock.Unlock()
	m.SelectionChan[sid] = c
}
