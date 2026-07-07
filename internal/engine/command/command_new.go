package command

import (
	"context"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/spec/chat"
)

// cmdNew deletes old session and creates a new session
func (h *Handler) cmdNew(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	if sessionId == "" {
		respMessage.Text = &chat.TextPayload{Text: i18n.T(keys.CmdSessionNotActive)}
		return instance.SendMessage(ctx, respMessage)
	}

	_, avail := h.sessionManager.ChatSession.GetSession(sessionId)
	if !avail {
		respMessage.Text = &chat.TextPayload{Text: i18n.T(keys.CmdSessionNotActive)}
		return instance.SendMessage(ctx, respMessage)
	}

	p, m, _ := h.sessionManager.ChatSession.GetModel(sessionId)

	h.sessionManager.ChatSession.DeleteSession(sessionId)

	err := h.chatService.DeleteChatSession(ctx, instance, msg, sessionId, true)
	if err != nil {
		return err
	}

	sid, err := h.chatService.NewChatSession(ctx, instance, msg.ChatID, p, m)
	if err != nil {
		return err
	}

	err = h.repo.ReplaceChatSessions(ctx, sessionId, sid)
	if err != nil {
		return err
	}

	respMessage.Text = &chat.TextPayload{
		Text: i18n.T(keys.CmdNewSuccess, map[string]any{
			"Old":   sessionId,
			"New":   sid,
			"Name":  p,
			"Model": m,
		}),
	}
	return instance.SendMessage(ctx, respMessage)
}
