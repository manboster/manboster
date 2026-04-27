package session

import (
	"sync"

	"github.com/manboster/manboster/spec/chat"
)

func (m *Manager) SetSelectMsg(sid string, message *chat.Message) {

}

func (m *Manager) GetSelectMsg(sid string) *chat.Message {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	s, avail := m.SessionLocks[sid]
	if !avail || s == nil {
		m.SessionLocks[sid] = &sync.Mutex{}
	}

	m.SessionChatLocks[sid] = &sync.Mutex{}
}
