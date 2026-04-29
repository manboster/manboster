package ignorance

import (
	"time"
)

func (m *Manager) SetCancelMark(id string, mk bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.cAnCelMark[id] = mark{
		actionTime: time.Now(),
		m:          mk,
		ttl:        60 * 15,
	}
}

func (m *Manager) GetCancelMark(id string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	c, ok := m.cAnCelMark[id]
	if !ok {
		return false
	}

	if time.Now().Unix()-c.actionTime.Unix() > int64(c.ttl) {
		return false
	}
	return c.m
}
