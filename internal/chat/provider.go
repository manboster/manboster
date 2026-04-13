package chat

import (
	"context"
)

// Provider defines which you want to implement, you can use Telegram, slack, even webserver api to chat with.
type Provider interface {
	Start(ctx context.Context, config any, handlerFunc func(msg *Message)) error
	SendMessage(ctx context.Context, msg *Message) error
	EditMessage(ctx context.Context, msg *Message) error
	Select(ctx context.Context, sessionId string, msg *Message) error // returned session id, If AbilityType & AbilitySelect == 0, this function is null and unable to send it.
	Stop(ctx context.Context) error
	Notify(ctx context.Context, chatID string, action ActionType) error
	Name() string
	New() Provider
	Ability() AbilityType // return the ability, which is this provider able to do.
}

// ActionType gives you the type of current action's callback.
type ActionType string

const (
	ActionUnknown ActionType = ""
	ActionTyping  ActionType = "typing"
)
