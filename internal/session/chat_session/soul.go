package chat_session

func (m *Manager) SetSoul(sid string, soul []string) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	s, avail := m.Sessions[sid]
	if !avail {
		s := Session{
			Souls: soul,
		}
		m.Sessions[sid] = s
		return
	}
	s.Souls = soul
	m.Sessions[sid] = s
}

func (m *Manager) GetSoul(sid string) ([]string, bool) {
	m.Lock.RLock()
	defer m.Lock.RUnlock()

	s, avail := m.Sessions[sid]
	if !avail {
		return s.Souls, false
	}
	return s.Souls, true
}

func (m *Manager) RemoveSoul(sid string, soul string) {
	m.Lock.Lock()
	defer m.Lock.Unlock()
	s, avail := m.Sessions[sid]

	if !avail {
		return
	}

	for i, _ := range s.Souls {
		if s.Souls[i] == soul {
			s.Souls = append(s.Souls[:i], s.Souls[i+1:]...)
			return
		}
	}
}

func (m *Manager) AppendSoul(sid string, soul string) {
	m.Lock.Lock()
	defer m.Lock.Unlock()
	s, avail := m.Sessions[sid]
	if !avail {
		s := Session{
			Souls: []string{soul},
		}
		m.Sessions[sid] = s
		return
	}
	for i, _ := range s.Souls {
		if s.Souls[i] == soul {
			return
		}
	}
	s.Souls = append(s.Souls, soul)
}
