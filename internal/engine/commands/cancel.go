package commands

import (
	"context"

	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/session"
)

// Cancel enables user to cancel their request
func Cancel(ctx context.Context, instance chat.Provider, msg *chat.Message, sessManager *session.Manager) error {
	msg.MessageType = chat.MessageText

	sessionId := sessManager.ID(instance.Name(), msg.ChatID)
	sessData, avail := sessManager.GetSession(sessionId)

	var text string
	if avail {
		sessData.Cancel()
		text = "[Manboster] Successfully cancelled the request."
	} else {
		text = "[Manboster] Failed to cancel the request: Session Object not found."
	}

	msg.Text = &chat.TextPayload{
		Text: text,
	}

	return instance.SendMessage(ctx, msg)
}
