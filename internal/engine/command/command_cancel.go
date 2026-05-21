package command

import (
	"context"

	"github.com/manboster/manboster/spec/chat"
)

// cmdCancel enables user to cancel their request
func (h *Handler) cmdCancel(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	msg.MessageType = chat.MessageText
	sessData, avail := h.sessionService.Manager.ChatSession.GetSession(sessionId)

	var text string
	if avail {
		if sessData.Active {
			sessData.Cancel()
			text = "[Manboster] Successfully cancelled the request."
		} else {
			text = "[Manboster] The request in this session is not active."
		}
	} else {
		text = "[Manboster] The request in this session is not active."
	}

	msg.Text = &chat.TextPayload{
		Text: text,
	}

	return instance.SendMessage(ctx, msg)
}
