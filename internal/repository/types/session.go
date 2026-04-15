package types

import (
	"time"

	"github.com/manboster/manboster/internal/database/types"
)

type Session struct {
	ID               uint64
	SessionID        string
	LLMProviderModel string
	LLMProvider      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func MapSess(session Session) types.Session {
	return types.Session{
		ID:               session.ID,
		SessionID:        session.SessionID,
		LLMProviderModel: session.LLMProviderModel,
		LLMProvider:      session.LLMProvider,
		CreatedAt:        session.CreatedAt,
		UpdatedAt:        session.UpdatedAt,
	}
}

func MapSession(session types.Session) Session {
	return Session{
		ID:               session.ID,
		SessionID:        session.SessionID,
		LLMProviderModel: session.LLMProviderModel,
		LLMProvider:      session.LLMProvider,
		CreatedAt:        session.CreatedAt,
		UpdatedAt:        session.UpdatedAt,
	}
}
