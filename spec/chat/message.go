package chat

import (
	"time"
)

// Message defines the universal definition of a provider's message standard.
type Message struct {
	Provider    string      // Required. Provider Platform like Telegram.
	MessageID   string      // Required. Global ID of a message, if this is request, it's the message id that replies.
	ChatID      string      // Required. The chat id
	UserID      string      // Required. Sender user's identify ID
	Username    string      // Required. Sender username
	MessageType MessageType // Required. reserved to stickers or more...
	ChatType    ChatsType   // Required. Chat's type, like Group, Channel, and Personal Chats.
	CreatedAt   time.Time   // Required. When is the message created?

	Reply   *Message // Optional. When it's valid, it means the message replied
	Forward *Message // Optional. When it's valid, it means the message was forwarded from someone.

	ChatName string // Optional. The chat's name, Required when ChatsType = ChatsGroup or ChatsChannel

	Command           *CommandPayload           // Optional. Required when MessageType = MessageCommand
	Selection         *SelectionPayload         // Optional. Required when MessageType = MessageSelection
	SelectionCallback *SelectionCallbackPayload // Optional. Required when MessageType = MessageSelectionCallback
	Text              *TextPayload              // Optional. Required when MessageType = MessageText
}

// MessageType is an enum defining msg types.
type MessageType int

const (
	MessageUnknown MessageType = 1 << iota
	MessageText
	MessageCommand
	MessageThinkingText
	MessageImage
	MessageVoice
	MessageVideo
	MessageFile
	MessageSelectionCallback
	MessageSelection
	MessageFromCron
	MessageFromCronIgnore
	MessageStart
	MessageFromRetry
)

const MessageTextAndImage = MessageText | MessageImage
const MessageTextImageAndFile = MessageTextAndImage | MessageFile

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
