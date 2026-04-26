package chat

import (
	"context"

	"github.com/manboster/manboster/internal/config"
)

// Provider defines which you want to implement, you can use Telegram, slack, even webserver api to chat with.
type Provider interface {
	Init(ctx context.Context, config any) error
	Start(ctx context.Context, handlerFunc func(msg *Message)) error
	SendMessage(ctx context.Context, msg *Message) error
	EditMessage(ctx context.Context, msg *Message) error
	Select(ctx context.Context, sessionId string, msg *Message) error // returned session id, If AbilityType & AbilitySelect == 0, this function is null and unable to send it.
	Stop() error
	Notify(ctx context.Context, msg *Message, action ActionType) error
	Name() string
	DisplayName() string
	New() Provider
	Ability() AbilityType // return the ability, which is this provider able to do.
	Config() config.Provider
}
