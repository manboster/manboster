package ignorance

import "time"

func (m *Manager) SetIgnoreMark(id string, mk bool, ttl int) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.ignoreMark[id] = mark{
		m:          mk,
		ttl:        ttl,
		actionTime: time.Now(),
	}
}

func (m *Manager) GetIgnoreMark(id string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	im, ok := m.ignoreMark[id]
	if !ok {
		return false
	}

	if time.Now().Unix()-im.actionTime.Unix() > int64(im.ttl) {
		return false
	}
	return im.m
}
