package engine

import (
	"context"

	"github.com/manboster/manboster/internal/chat"
)

// HandleReject plays a reject role of the application
func (e *Engine) HandleReject(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	msg.MessageType = chat.MessageText
	msg.Text = &chat.TextPayload{
		Text: "Access denied. You are not allowed to use this bot or this command.\nIf you are sure that this is not an error, please contact this bot's owner.",
	}
	return instance.SendMessage(ctx, msg)
}
