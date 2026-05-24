package handler

import (
	"context"
	"strings"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/spec/chat"
)

// HandleOnBoard helps user tackle onboarding problems
func (h *Handler) HandleOnBoard(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	if h.onboard.Active() {
		h.onboard.HandleOnBoard()
	}
	respMessage := msg.Clone()

	var text strings.Builder
	text.WriteString(i18n.T(keys.EngineOnboardWelcome))
	text.WriteString(i18n.T(keys.EngineOnboardInstruction))
	text.WriteString(i18n.T(keys.EngineOnboardStep1))
	text.WriteString(i18n.T(keys.EngineOnboardStep1Note))
	text.WriteString(i18n.T(keys.EngineOnboardStep2))
	text.WriteString(i18n.T(keys.EngineOnboardStep3))
	text.WriteString(i18n.T(keys.EngineOnboardWish))

	respMessage.MessageType = chat.MessageText
	respMessage.Text = &chat.TextPayload{
		Text: text.String(),
	}
	return instance.SendMessage(ctx, respMessage)
}
