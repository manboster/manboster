package command

import (
	"context"

	"github.com/manboster/manboster/spec/chat"
)

func (h *Handler) cmdDefault(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	msg.MessageType = chat.MessageText
	msg.Text = &chat.TextPayload{
		Text: "We are sorry but this is an Invalid Command. Please check your grammatical mistakes.",
	}
	return instance.SendMessage(ctx, msg)
}
