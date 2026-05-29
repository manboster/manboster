package session

import (
	"context"

	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/internal/session/chat_session"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/llm"
)

func (s *Service) NewChatSession(ctx context.Context, provider string, llmProvider string, model string, chatId string) (string, error) {
	sessionId := util.RandomString(8)
	soulsList := s.soulService.GetSoulsList(ctx, chatId)
	err := s.repo.CreateSession(ctx, types.Session{
		SessionID:        sessionId,
		LLMProvider:      llmProvider,
		LLMProviderModel: model,
		ActivatedSouls:   soulsList,
	})
	if err != nil {
		return "", err
	}
	// set a new session
	s.Manager.ChatSession.SetSession(sessionId, chat_session.Session{
		Model:    model,
		Provider: llmProvider,
		Events:   []llm.Event{},
		Souls:    soulsList,
		Active:   false,
		Cancel:   nil,
	})

	_, err = s.repo.GetChat(ctx, chatId, provider)
	if err != nil {
		err = s.repo.CreateChat(ctx, types.Chat{
			ChatID:         chatId,
			SessionID:      sessionId,
			ChatProvider:   provider,
			ChatPermission: 1, // TODO: add chat permission but there is no need, so occupy?
		})

	} else {
		err = s.repo.UpdateChat(ctx, chatId, provider, sessionId)
	}
	if err != nil {
		return "", err
	}

	return sessionId, nil
}
