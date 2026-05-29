package chat_session

import "github.com/manboster/manboster/spec/chat"

func (m *Manager) GetInputMsgID(sid string) string {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	s, ok := m.Sessions[sid]
	if !ok {
		return ""
	}

	return s.InputMsgID
}

func (m *Manager) GetInputMsg(sid string) *chat.Message {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	s, ok := m.Sessions[sid]
	if !ok {
		return nil
	}
	return s.InputMsg
}

func (m *Manager) SetMsg(sid string, msg *chat.Message) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	s, ok := m.Sessions[sid]
	if !ok {
		s = Session{
			InputMsg: msg,
		}
		m.Sessions[sid] = s
		return
	}

	s.InputMsg = msg
	s.InputMsgID = msg.MessageID
	m.Sessions[sid] = s
}

func (m *Manager) SetMsgId(sid string, id string) {
	m.Lock.Lock()
	defer m.Lock.Unlock()

	s, ok := m.Sessions[sid]
	if !ok {
		s = Session{
			InputMsgID: id,
		}
		m.Sessions[sid] = s
		return
	}

	s.InputMsgID = id
	m.Sessions[sid] = s
}
