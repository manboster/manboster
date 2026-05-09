package ignorance

func (m *Manager) SetHachimiCache(params string, user bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.hachimiMark[params] = user
}

func (m *Manager) GetHachimiCache(params string) (bool, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	u, ok := m.hachimiMark[params]
	return u, ok
}
