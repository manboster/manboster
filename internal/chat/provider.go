package chat

import (
	"context"
)

// Provider defines which you want to implement, you can use Telegram, slack, even webserver api to chat with.
type Provider interface {
	Start(ctx context.Context, config any, handlerFunc func(msg *Message)) error
	SendMessage(ctx context.Context, msg *Message) error
	EditMessage(ctx context.Context, msg *Message) error
	Select(ctx context.Context, sessionId string, msg *Message) error // returned session id
	Stop(ctx context.Context) error
	Notify(chatID string, action ActionType) error
	Name() string
	New() Provider
}

// ActionType gives you the type of current action's callback.
type ActionType string

const (
	ActionUnknown ActionType = ""
	ActionTyping  ActionType = "typing"
)
