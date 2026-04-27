package session

import (
	"github.com/manboster/manboster/internal/session/chat"
	"github.com/manboster/manboster/internal/session/chat_session"
	"github.com/manboster/manboster/internal/session/ignorance"
	"github.com/manboster/manboster/internal/session/selection"
)

type Manager struct {
	ChatSession      *chat_session.Manager
	Chat             *chat.Manager
	SelectionManager *selection.Manager
	Ignorance        *ignorance.Manager
}

func NewManager() *Manager {
	return &Manager{
		ChatSession:      chat_session.New(),
		Chat:             chat.New(),
		SelectionManager: selection.New(),
		Ignorance:        ignorance.New(),
	}
}
