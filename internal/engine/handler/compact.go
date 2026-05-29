package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/engine/chatdata"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

// HandleCompact compacts data
func (h *Handler) HandleCompact(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) (string, error) {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	text := ""
	respMessage.Text = &chat.TextPayload{
		Text: i18n.Te(keys.EngineHandlerCompactWait, sessionId, nil),
	}
	err := instance.SendMessage(ctx, respMessage)
	if err != nil {
		return "", err
	}

	newSession, err := h.chatDataService.Compact(ctx, instance, msg, sessionId)
	if err != nil {
		if errors.Is(err, chatdata.ErrNoNeedToCompact) {
			text = i18n.T(keys.EngineHandlerCompactNoNeed)
		} else {
			color.Yellow(fmt.Sprintf("[Manboster Handler] We encountered an error when compacting message: %q", err))
			text = i18n.T(keys.EngineHandlerCompactError)
		}
		respMessage.MessageType = chat.MessageText
		respMessage.Text = &chat.TextPayload{
			Text: text,
		}
		err = instance.SendMessage(ctx, respMessage)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Handler] We encountered an error when sending message: %q", err))
		}
		return "", err
	}
	return newSession, nil
}

// CheckCompact returns value about whether this should compact or not
func (h *Handler) CheckCompact(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) (bool, string, error) {
	provider, model, _ := h.sessionManager.ChatSession.GetModel(sessionId)
	_, m := util.GetModelWithFallback(ctx, h.llmProviders, provider, model)

	// get total tokens in order to compact
	totToken, err := h.repo.GetTotalToken(ctx, sessionId)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] Error while getting total tokens from repository: %q", err))
		return false, "", nil
	}
	// checkout whether a need to compact or not
	if uint64(totToken) > llm.CalculateCompactTokens(m) {
		newSession, err := h.HandleCompact(ctx, instance, msg, sessionId)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] Error while compacting data: %q", err))
			return false, "", err
		}
		return true, newSession, err
	}
	return false, "", nil
}
