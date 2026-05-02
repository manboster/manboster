package session

import (
	"context"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/spec/chat"
)

// LoadChatSession helps Manboster Engine get sessionId, preparing for the message handler
func (s *Service) LoadChatSession(ctx context.Context, instance chat.Provider, msg *chat.Message, isAdmin bool) (string, error) {
	lockerID := fmt.Sprintf("%s:%s", instance.Name(), msg.ChatID)
	chatLock := s.Manager.Chat.GetSessionChatLocks(lockerID)

	var sessionId string
	chatLock.Lock()
	defer chatLock.Unlock()

	chatInfo, err := s.repo.GetChat(ctx, msg.ChatID, instance.Name())
	if err == nil {
		sessionId = chatInfo.SessionID
		// we need to read model and provider.
		sessInfo, err := s.repo.GetSession(ctx, sessionId)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Session Service] We encountered an error while getting chat data, error: %q", err))
			return "", err
		}
		s.Manager.ChatSession.SetModel(sessInfo.SessionID, sessInfo.LLMProvider, sessInfo.LLMProviderModel)
		s.Manager.ChatSession.SetSoul(sessionId, sessInfo.ActivatedSouls)
	} else if errors.Is(err, repository.ErrNotFound) {
		// if you're not an administrator, you can not create a new session
		if isAdmin {
			sid, err := s.NewChatSession(ctx, instance.Name(), msg)
			sessionId = sid
			if err != nil {
				color.Red(fmt.Sprintf("[Manboster Session Service] We encountered an error while creating session to repository, error: %q", err))
				return "", err
			}
		} else {
			// return access denied and reject it
			return "", ErrAccessDenied
		}
	} else {
		color.Red(fmt.Sprintf("[Manboster Session Service] We encountered an error while reading user information to repository, error: %q", err))
		return "", err
	}

	err = s.MergeChatSession(ctx, sessionId)
	if err != nil {
		return "", err
	}

	color.Blue("[Manboster Engine] This session is not available in memory storage, now loading from database")

	return sessionId, nil
}
