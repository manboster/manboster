package commands

import (
	"context"

	"github.com/manboster/manboster/internal/chat"
)

func Default(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	msg.MessageType = chat.MessageText
	msg.Text = &chat.TextPayload{
		Text: "We are sorry but this is an Invalid Command. Please check your grammatical mistakes.",
	}
	return instance.SendMessage(ctx, msg)
}
