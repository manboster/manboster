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
	Provider    string      // Required. Provider Platform like Telegram.
	MessageID   string      // Required. Global ID of a message, if this is request, it's the message id that replies.
	ChatID      string      // Required. The chat id
	UserID      string      // Required. Sender user's identify ID
	Username    string      // Required. Sender username
	MessageType MessageType // Required. reserved to stickers or more...
	ChatType    ChatsType   // Required. Chat's type, like Group, Channel, and Personal Chats.

	Reply *Message // Optional. When it's valid, it means the message replied

	Text string // Optional. Required when MessageType = MessageText Text Info

	CommandType CommandType // Optional. Required when MessageType = MessageCommand Command's type
	CommandArgs []string    // Optional. Required when MessageType = MessageCommand Command's args

	SelectionSession string // Optional. Required when MessageType == MessageSelectionCallback, it should have a value.
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

// ActionType gives you the type of current action's callback.
type ActionType string

const (
	ActionUnknown ActionType = ""
	ActionTyping  ActionType = "typing"
)

// CommandType defines command's types.
type CommandType string

const (
	CommandUnknown CommandType = ""        // No Command Available
	CommandVersion CommandType = "version" // Get Version Data
	CommandHelp    CommandType = "help"    // Get Helper Messages
	CommandOp      CommandType = "op"      // Grant a user
	CommandDeOp    CommandType = "deop"    // Ungrant a user
	CommandId      CommandType = "id"      // display ids
	CommandStatus  CommandType = "status"  // display current status
	CommandSave    CommandType = "save"    // save this chat to database
	CommandNew     CommandType = "new"     // delete this and create a new chat
	CommandSummary CommandType = "summary" // summary this chat and create a new chat with summarized items
	CommandModels  CommandType = "models"  // select models you want
	CommandStart   CommandType = "start"   // start command gives tips to you, when it's the first run, it will grant you the root access to this application.
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
