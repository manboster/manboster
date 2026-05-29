package chat

func (m *Message) Fork() *Message {
	if m == nil {
		return nil
	}

	return &Message{
		Provider:    m.Provider,
		MessageID:   m.MessageID,
		ChatID:      m.ChatID,
		UserID:      m.UserID,
		Username:    m.Username,
		MessageType: m.MessageType,
		ChatType:    m.ChatType,
		CreatedAt:   m.CreatedAt,

		Reply:   m.Reply,
		Forward: m.Forward,

		ChatName: m.ChatName,

		Command:           m.Command,
		Selection:         m.Selection,
		SelectionCallback: m.SelectionCallback,
		Text:              m.Text,
	}
}
