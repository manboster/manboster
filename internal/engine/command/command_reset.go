package command

import (
	"context"

	"github.com/manboster/manboster/internal/engine/gatekeeper"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/spec/chat"
)

// cmdReset resets all memories of approve or disallow in this session and chat
func (h *Handler) cmdReset(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText

	if sessionId == "" {
		respMessage.Text = &chat.TextPayload{
			Text: i18n.T(keys.CmdSessionNotActive),
		}

		return instance.SendMessage(ctx, respMessage)
	}

	h.sessionService.Manager.Ignorance.Clear(gatekeeper.BuildSessionId(instance.Name(), msg.ChatID, sessionId))
	respMessage.Text = &chat.TextPayload{
		Text: i18n.T(keys.CmdResetSuccess),
	}
	return instance.SendMessage(ctx, respMessage)
}
