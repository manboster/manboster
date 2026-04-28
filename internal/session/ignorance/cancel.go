package ignorance

import (
	"time"
)

func (m *Manager) SetCancelMark(id string, mark bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.cAnCelMark[id] = cAnCel{
		actionTime: time.Now(),
		isCancel:   mark,
	}
}

func (m *Manager) GetCancelMark(id string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	c, ok := m.cAnCelMark[id]
	if !ok {
		return false
	}

	if time.Now().Unix()-c.actionTime.Unix() > 15*60 {
		return false
	}
	return c.isCancel
}
