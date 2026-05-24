package command

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/spec/chat"
)

// cmdSave saves old session and create a new session
func (h *Handler) cmdSave(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText

	if sessionId == "" {
		respMessage.Text = &chat.TextPayload{Text: i18n.T(keys.CmdSessionNotActive)}
		return instance.SendMessage(ctx, respMessage)
	}

	_, avail := h.sessionService.Manager.ChatSession.GetSession(sessionId)
	if !avail {
		respMessage.Text = &chat.TextPayload{Text: i18n.T(keys.CmdSessionNotActive)}
		return instance.SendMessage(ctx, respMessage)
	}

	p, m, _ := h.sessionService.Manager.ChatSession.GetModel(sessionId)
	h.sessionService.Manager.ChatSession.DeleteSession(sessionId)
	err := h.repo.DeleteChat(ctx, msg.ChatID, instance.Name())
	if err != nil {
		return err
	}

	sid, err := h.sessionService.NewChatSession(ctx, instance.Name(), p, m, msg)
	if err != nil {
		return err
	}

	respMessage.Text = &chat.TextPayload{
		Text: fmt.Sprintf(i18n.T(keys.CmdSaveSuccess), sessionId, sid, p, m),
	}
	return instance.SendMessage(ctx, respMessage)
}
