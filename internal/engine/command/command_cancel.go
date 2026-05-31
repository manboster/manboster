package command

import (
	"context"

	"github.com/manboster/manboster/spec/chat"
)

// cmdCancel enables user to cancel their request
func (h *Handler) cmdCancel(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	msg.MessageType = chat.MessageText
	sessData, avail := h.sessionService.Manager.ChatSession.GetSession(sessionId)

	isDisplay := false
	if val, ok := ctx.Value("d").(bool); ok {
		isDisplay = val
	}

	var text string
	if avail {
		if sessData.Active {
			h.sessionService.Manager.ChatSession.Deactivate(sessionId)
			h.sessionService.Manager.ChatSession.SessionCancel(sessionId)
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

	if isDisplay {
		return instance.SendMessage(ctx, msg)
	}
	return nil
}
