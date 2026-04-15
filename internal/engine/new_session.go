package engine

import (
	"context"

	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/internal/session"
	"github.com/manboster/manboster/internal/util"
)

func (e *Engine) newSession(ctx context.Context, msg *chat.Message, provider string) (string, error) {
	sessionId := util.RandomString(32)
	err := e.repo.CreateSession(ctx, types.Session{
		SessionID:        sessionId,
		LLMProvider:      e.config.App.DefaultLLMProvider,
		LLMProviderModel: e.config.App.DefaultLLMModel,
	})
	if err != nil {
		return "", err
	}
	// set a new session
	e.sessionManager.SetSession(sessionId, session.Session{
		Model:    e.config.App.DefaultLLMModel,
		Provider: e.config.App.DefaultLLMProvider,
		Events:   []llm.Event{},
		Active:   false,
		Cancel:   nil,
	})

	err = e.repo.CreateChat(ctx, types.Chat{
		ChatID:         msg.ChatID,
		SessionID:      sessionId,
		ChatProvider:   provider,
		ChatPermission: 1, // TODO: add chat permission but there is no need, so occupy?
	})
	if err != nil {
		return "", err
	}
	return sessionId, nil
}
