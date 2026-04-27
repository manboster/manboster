package ignorance

func (m *Manager) SetIgnoreMark(id string, mark bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.ignoreMark[id] = mark
}

func (m *Manager) GetIgnoreMark(id string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	im, ok := m.ignoreMark[id]
	if !ok {
		return false
	}
	return im
}
