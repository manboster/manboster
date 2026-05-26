package chat

func (m *Manager) GetToolMsgData(sid string) string {
	m.Lock.RLock()
	defer m.Lock.RUnlock()

	val, ok := m.SessionChatData[sid]
	if !ok {
		return ""
	}
	return val.MsgData
}

func (m *Manager) GetToolMsgId(sid string) string {
	m.Lock.RLock()
	defer m.Lock.RUnlock()
	val, ok := m.SessionChatData[sid]
	if !ok {
		return ""
	}
	val.Counter++
	m.SessionChatData[sid] = val
	return val.MsgId
}

func (m *Manager) GetToolCallCounts(sid string) int {
	m.Lock.RLock()
	defer m.Lock.RUnlock()
	val, ok := m.SessionChatData[sid]
	if !ok {
		return 0
	}
	return val.Counter
}

func (m *Manager) ResetTool(sid string, msgId string) {
	m.Lock.Lock()
	defer m.Lock.Unlock()
	val := data{
		MsgId:   msgId,
		Counter: 1,
	}
	m.SessionChatData[sid] = val
}

func (m *Manager) SetToolMsgData(sid string, data string) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	val, ok := m.SessionChatData[sid]
	if !ok {
		return
	}
	val.MsgData = data
	m.SessionChatData[sid] = val
}
