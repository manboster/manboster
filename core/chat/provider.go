package chat

import "context"

// Provider defines which you want to implement, you can use Telegram, slack, even webserver api to chat with.
type Provider interface {
	Start(ctx context.Context, handlerFunc func(msg *Message)) error
	SendMessage(ctx context.Context, msg Message) error
	Select(ctx context.Context, title string, name string, selection []Selection) (string, error) // returned session id
	Stop(ctx context.Context) error
}

// Message defines the universal definition of a provider's message standard.
type Message struct {
	Provider         string      // Provider Platform like Telegram.
	MessageID        string      // Global ID of a message
	ChatID           string      // The chat id
	UserID           string      // Sender user's identify ID
	Username         string      // Sender username
	MessageType      MessageType // TODO: reserved to stickers or more... Now we define 1 is Text.
	ChatType         ChatsType   // Chat's type, like Group, Channel, and Personal Chats.
	Text             string      // Text Info
	SelectionSession string      // if MessageType == MessageTypeSelectionCallback, it should have a value.
}

// MessageType is an enum defining msg types.
type MessageType int16

const (
	MessageTypeUnknown           MessageType = 0
	MessageTypeText              MessageType = 1
	MessageTypeSelectionCallback MessageType = 256
)

// ChatsType is an enum defining chat types.
type ChatsType int16

const (
	ChatsTypeUnknown  ChatsType = 0
	ChatsTypePersonal ChatsType = 1
	ChatsTypeGroup    ChatsType = 2
	ChatsTypeChannel  ChatsType = 3
)

// Selection provides options to select, answer should be the value.
type Selection struct {
	Name  string // display name
	Value string // actual value
}
