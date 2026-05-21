package command

import (
	"context"

	"github.com/manboster/manboster/spec/chat"
)

type Command struct {
	DisplayName string
	Name        string
	Description string
}

type handleFunc func(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error

type HandlerInterface[T ~string] interface {
	Register(t T, fn handleFunc)
	Handle(ctx context.Context, t T) error
}
