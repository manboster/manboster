package types

import (
	"encoding/json"
	"time"

	"github.com/manboster/manboster/internal/database/types"
)

type Session struct {
	ID               uint64
	SessionID        string
	LLMProviderModel string
	LLMProvider      string
	ActivatedSouls   []string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (session Session) Map() types.Session {
	jsonify, _ := json.Marshal(session.ActivatedSouls)
	return types.Session{
		ID:               session.ID,
		SessionID:        session.SessionID,
		LLMProviderModel: session.LLMProviderModel,
		LLMProvider:      session.LLMProvider,
		ActivatedSouls:   string(jsonify),
		CreatedAt:        session.CreatedAt,
		UpdatedAt:        session.UpdatedAt,
	}
}

func SessionFrom(session types.Session) Session {
	var activatedSouls []string
	_ = json.Unmarshal([]byte(session.ActivatedSouls), &activatedSouls)

	return Session{
		ID:               session.ID,
		SessionID:        session.SessionID,
		LLMProviderModel: session.LLMProviderModel,
		LLMProvider:      session.LLMProvider,
		ActivatedSouls:   activatedSouls,
		CreatedAt:        session.CreatedAt,
		UpdatedAt:        session.UpdatedAt,
	}
}
