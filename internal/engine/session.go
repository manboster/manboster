package engine

import (
	"context"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/repository"
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

// loadSession helps Manboster Engine get sessionId, preparing for the message handler
func (e *Engine) loadSession(ctx context.Context, instance chat.Provider, msg *chat.Message, isAdmin bool) (string, error) {
	lockerID := fmt.Sprintf("%s:%s", instance.Name(), msg.ChatID)
	chatLock := e.sessionManager.GetSessionChatLocks(lockerID)

	var sessionId string
	chatLock.Lock()
	defer chatLock.Unlock()

	chatInfo, err := e.repo.GetChat(ctx, msg.ChatID, instance.Name())
	if err == nil {
		sessionId = chatInfo.SessionID
	} else if errors.Is(err, repository.ErrNotFound) {
		// if you're not an administrator, you can not create a new session
		if isAdmin {
			sid, err := e.newSession(ctx, msg, instance.Name())
			sessionId = sid
			if err != nil {
				color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while creating session to repository, error: %q", err))
				return "", err
			}
		} else {
			// return access denied and reject it
			return "", ErrAccessDenied
		}
	} else {
		color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while reading user information to repository, error: %q", err))
		return "", err
	}

	return sessionId, nil
}
