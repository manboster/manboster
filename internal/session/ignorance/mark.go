package ignorance

import "time"

func (m *Manager) SetMark(id string, mk bool, ttl int, mType MarkType) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.mark[id] = mark{
		M:          mk,
		Ttl:        ttl,
		ActionTime: time.Now(),
		MarkType:   mType,
	}
}

// UpdateMark updates current last updated time
func (m *Manager) UpdateMark(id string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	im, ok := m.mark[id]
	if !ok {
		return
	}
	im.ActionTime = time.Now()
	m.mark[id] = im
}

func (m *Manager) GetMark(id string) (bool, MarkType) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	im, ok := m.mark[id]
	if !ok {
		return false, ""
	}

	if time.Now().Unix()-im.ActionTime.Unix() > int64(im.Ttl) {
		return false, im.MarkType
	}
	return im.M, im.MarkType
}
