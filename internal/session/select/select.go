package selection

import (
	"github.com/manboster/manboster/spec/chat"
)

func (m *Manager) SetSelectMsg(sid string, message *chat.Message) {
	m.SelectionLock.Lock()
	defer m.SelectionLock.Unlock()
	m.Selection[sid] = message
}

func (m *Manager) GetSelectMsg(sid string) *chat.Message {
	m.SelectionLock.Lock()
	defer m.SelectionLock.Unlock()
	s, ok := m.Selection[sid]
	if !ok {
		return nil
	}
	return s
}

func (m *Manager) CleanSelect(sid string) {
	m.SelectionLock.Lock()
	defer m.SelectionLock.Unlock()
	delete(m.Selection, sid)
	delete(m.SelectionChan, sid)
}
