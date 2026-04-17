package commands

import (
	"context"

	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/session"
)

// Cancel enables user to cancel their request
func Cancel(ctx context.Context, instance chat.Provider, msg *chat.Message, sessManager *session.Manager, sessionId string) error {
	msg.MessageType = chat.MessageText
	sessData, avail := sessManager.GetSession(sessionId)

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
