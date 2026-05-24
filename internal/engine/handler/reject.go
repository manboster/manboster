package handler

import (
	"context"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/spec/chat"
)

// HandleReject plays a reject role of the application
func (h *Handler) HandleReject(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	msg.MessageType = chat.MessageText
	msg.Text = &chat.TextPayload{
		Text: i18n.T(keys.EngineHandlerReject),
	}
	return instance.SendMessage(ctx, msg)
}
