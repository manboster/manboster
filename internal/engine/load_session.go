package engine

import (
	"context"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/repository"
)

func (e *Engine) loadSession(ctx context.Context, instance chat.Provider, msg *chat.Message) (string, error) {
	lockerID := fmt.Sprintf("%s:%s", instance.Name(), msg.ChatID)
	chatLock := e.sessionManager.GetSessionChatLocks(lockerID)

	var sessionId string
	chatLock.Lock()
	defer chatLock.Unlock()

	chatInfo, err := e.repo.GetChat(ctx, msg.ChatID, instance.Name())
	if err == nil {
		sessionId = chatInfo.SessionID
	} else if errors.Is(err, repository.ErrNotFound) {
		sid, err := e.newSession(ctx, msg, instance.Name())
		sessionId = sid
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while creating session to repository, error: %q", err))
			return "", err
		}
	} else {
		color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while reading user information to repository, error: %q", err))
		return "", err
	}

	return sessionId, nil
}
