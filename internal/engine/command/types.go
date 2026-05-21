package command

import (
	"context"
)

type Command struct {
	DisplayName string
	Name        string
	Description string
}

type handleFunc func(ctx context.Context) error

type HandlerInterface[T ~string] interface {
	Register(t T, fn handleFunc)
	Handle(ctx context.Context, t T) error
}
