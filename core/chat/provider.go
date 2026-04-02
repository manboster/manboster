package chat

import (
	"context"
)

// Provider defines which you want to implement, you can use Telegram, slack, even webserver api to chat with.
type Provider interface {
	Start(ctx context.Context, config any, handlerFunc func(msg *Message)) error
	SendMessage(ctx context.Context, msg *Message) error
	Select(ctx context.Context, title string, name string, selection []Selection) (string, error) // returned session id
	Stop(ctx context.Context) error
	Notify(chatID string, action ActionType) error
	Name() string
	New() Provider
}

// Message defines the universal definition of a provider's message standard.
type Message struct {
	Provider         string      // Provider Platform like Telegram.
	MessageID        string      // Global ID of a message, if this is request, it's the message id that replies.
	ChatID           string      // The chat id
	UserID           string      // Sender user's identify ID
	Username         string      // Sender username
	MessageType      MessageType // TODO: reserved to stickers or more... Now we define 1 is Text.
	ChatType         ChatsType   // Chat's type, like Group, Channel, and Personal Chats.
	Text             string      // Text Info
	SelectionSession string      // if MessageType == MessageSelectionCallback, it should have a value.
}

// MessageType is an enum defining msg types.
type MessageType int16

const (
	MessageUnknown           MessageType = 0
	MessageText              MessageType = 1
	MessageCommand           MessageType = 255
	MessageSelectionCallback MessageType = 256
)

// ChatsType is an enum defining chat types.
type ChatsType int16

const (
	ChatsUnknown  ChatsType = 0
	ChatsPersonal ChatsType = 1
	ChatsGroup    ChatsType = 2
	ChatsChannel  ChatsType = 3
)

// Selection provides options to select, answer should be the value.
type Selection struct {
	Name  string // display name
	Value string // actual value
}

// ActionType gives you the type of current actions callback.
type ActionType string

const (
	ActionUnknown ActionType = ""
	ActionTyping  ActionType = "typing"
)

// CommandType defines command's types.
type CommandType string

const (
	CommandUnknown CommandType = ""
	CommandVersion CommandType = "version"
	CommandHelp    CommandType = "help"
	CommandGrant   CommandType = "grant"
)
