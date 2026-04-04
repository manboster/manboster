package chat

// Message defines the universal definition of a provider's message standard.
type Message struct {
	Provider    string      // Required. Provider Platform like Telegram.
	MessageID   string      // Required. Global ID of a message, if this is request, it's the message id that replies.
	ChatID      string      // Required. The chat id
	UserID      string      // Required. Sender user's identify ID
	Username    string      // Required. Sender username
	MessageType MessageType // Required. reserved to stickers or more...
	ChatType    ChatsType   // Required. Chat's type, like Group, Channel, and Personal Chats.

	Reply *Message // Optional. When it's valid, it means the message replied

	Command           *CommandPayload           // Optional. Required when MessageType = MessageCommand
	Selection         *SelectionPayload         // Optional. Required when MessageType = MessageSelection
	SelectionCallback *SelectionCallbackPayload // Optional. Required when MessageType = MessageSelectionCallback
	Text              *TextPayload              // Optional. Required when MessageType = MessageText
}

// MessageType is an enum defining msg types.
type MessageType int16

const (
	MessageUnknown           MessageType = 0
	MessageText              MessageType = 1
	MessageCommand           MessageType = 255
	MessageSelectionCallback MessageType = 256
	MessageSelection         MessageType = 257
)

// Build gives a builder of Message
func (*Message) Build(provider Provider) *Message {
	return &Message{
		Provider:    provider.Name(),
		MessageID:   "",
		ChatID:      "",
		UserID:      "",
		Username:    "",
		MessageType: MessageUnknown,
		ChatType:    ChatsUnknown,
	}
}
