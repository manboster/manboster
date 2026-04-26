package chat

func (m *Message) Clone() *Message {
	if m == nil {
		return &Message{}
	}
	return &Message{
		Provider:    m.Provider,
		MessageType: MessageUnknown,
		ChatID:      m.ChatID,
		UserID:      m.UserID,
		Username:    m.Username,
		ChatType:    m.ChatType,
		Reply:       m,
	}
}
