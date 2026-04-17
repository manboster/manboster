package engine

import (
	"context"
	"fmt"

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
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	_, avail := e.sessionManager.GetSession(sessionId)
	if !avail {
		respMessage.Text = &chat.TextPayload{
			Text: "Session is not active, there is nothing to do!",
		}
		return instance.SendMessage(ctx, respMessage)
	}

	e.sessionManager.DeleteSession(sessionId)
	err := e.repo.DeleteChat(ctx, msg.ChatID, instance.Name())
	if err != nil {
		return err
	}

	sid, err := e.loadSession(ctx, instance, msg, true)
	if err != nil {
		return err
	}

	respMessage.Text = &chat.TextPayload{
		Text: fmt.Sprintf("Old session %s reserved. New session: %s", sessionId, sid),
	}
	return instance.SendMessage(ctx, respMessage)
}
