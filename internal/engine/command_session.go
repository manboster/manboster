package engine

import (
	"context"

	"github.com/manboster/manboster/internal/chat"
)

// cmdCancel enables user to cancel their request
func (e *Engine) cmdCancel(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	msg.MessageType = chat.MessageText
	sessData, avail := e.sessionManager.GetSession(sessionId)

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

func (e *Engine) cmdNew(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	msg.MessageType = chat.MessageText
	_, avail := e.sessionManager.GetSession(sessionId)
	if !avail {
		msg.Text = &chat.TextPayload{
			Text: "Session is not active, there is nothing to do!",
		}
		return instance.SendMessage(ctx, msg)
	}

	e.sessionManager.DeleteSession(sessionId)
	sessionId, err := e.loadSession(ctx, instance, msg, true)
	if err != nil {
		return err
	}

	return instance.SendMessage(ctx, msg)
}
