package types

import "github.com/manboster/manboster/internal/database/types"

type Session struct {
	ID               uint64
	SessionID        string
	LLMProviderModel string
	LLMProvider      string
}

func MapSess(session Session) types.Session {
	return types.Session{
		ID:               session.ID,
		SessionID:        session.SessionID,
		LLMProviderModel: session.LLMProviderModel,
		LLMProvider:      session.LLMProvider,
	}
}

func MapSession(session types.Session) Session {
	return Session{
		ID:               session.ID,
		SessionID:        session.SessionID,
		LLMProviderModel: session.LLMProviderModel,
		LLMProvider:      session.LLMProvider,
	}
}
