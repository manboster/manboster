package chat

import (
	"context"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/spec/chat"
)

// LoadChatSession loads necessary chat session data into application
func (s *Service) LoadChatSession(ctx context.Context, instance chat.Provider, msg *chat.Message) (string, error) {
	lockerID := fmt.Sprintf("%s:%s", instance.Name(), msg.ChatID)
	chatLock := s.sessionManager.Chat.GetSessionChatLocks(lockerID)

	var sessionId string
	chatLock.Lock()
	defer chatLock.Unlock()

	chatInfo, err := s.repo.GetChat(ctx, msg.ChatID, instance.Name())
	if err == nil {
		sessionId = chatInfo.SessionID
		// we need to read model and provider from repository.
		sessInfo, err := s.repo.GetSession(ctx, sessionId)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Session Service] We encountered an error while getting chat data, error: %q", err))
			return "", err
		}

		s.sessionManager.ChatSession.SetModel(sessInfo.SessionID, sessInfo.LLMProvider, sessInfo.LLMProviderModel)
		s.sessionManager.ChatSession.SetSoul(sessionId, sessInfo.ActivatedSouls)

	} else if errors.Is(err, repository.ErrNotFound) {
		return "", ErrCreateRequired
	}

	color.Red(fmt.Sprintf("[Manboster Session Service] We encountered an error while reading user information to repository, error: %q", err))
	return "", err
}
