package ignorance

import "strings"

func (m *Manager) Clear(prefix string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	for key := range m.mark {
		if strings.HasPrefix(key, prefix) {
			delete(m.mark, key)
		}
	}
}
