package session

import (
	"context"

	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/internal/session/chat_session"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

func (s *Service) NewChatSession(ctx context.Context, provider string, msg *chat.Message) (string, error) {
	sessionId := util.RandomString(8)
	soulsList := s.soulService.GetSoulsList(ctx, msg.ChatID)
	err := s.repo.CreateSession(ctx, types.Session{
		SessionID:        sessionId,
		LLMProvider:      s.config.App.DefaultLLMProvider,
		LLMProviderModel: s.config.App.DefaultLLMModel,
		ActivatedSouls:   soulsList,
	})
	if err != nil {
		return "", err
	}
	// set a new session
	s.Manager.ChatSession.SetSession(sessionId, chat_session.Session{
		Model:    s.config.App.DefaultLLMModel,
		Provider: s.config.App.DefaultLLMProvider,
		Events:   []llm.Event{},
		Souls:    soulsList,
		Active:   false,
		Cancel:   nil,
	})

	err = s.repo.CreateChat(ctx, types.Chat{
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
