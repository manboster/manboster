package types

import (
	"time"

	"github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/spec/llm"
)

type EventData struct {
	ID               uint64
	EventID          string
	Role             llm.RoleType
	MessageType      llm.MessageType
	Model            string
	Provider         string
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
	MessagePayload   string // json encoded
	InputCost        float64
	OutputCost       float64
	CachedCost       float64
	TotalCost        float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func MapCD(eventData EventData) types.EventData {
	return types.EventData{
		ID:               eventData.ID,
		EventID:          eventData.EventID,
		Role:             string(eventData.Role),
		MessageType:      int16(eventData.MessageType),
		Model:            eventData.Model,
		Provider:         eventData.Provider,
		PromptTokens:     eventData.PromptTokens,
		CompletionTokens: eventData.CompletionTokens,
		TotalTokens:      eventData.TotalTokens,
		MessagePayload:   eventData.MessagePayload,
		InputCost:        eventData.InputCost,
		OutputCost:       eventData.OutputCost,
		TotalCost:        eventData.TotalCost,
		CreatedAt:        eventData.CreatedAt,
		UpdatedAt:        eventData.UpdatedAt,
	}
}

func MapEventData(eventData types.EventData) EventData {
	return EventData{
		ID:               eventData.ID,
		EventID:          eventData.EventID,
		Role:             llm.RoleType(eventData.Role),
		MessageType:      llm.MessageType(eventData.MessageType),
		Model:            eventData.Model,
		Provider:         eventData.Provider,
		PromptTokens:     eventData.PromptTokens,
		CompletionTokens: eventData.CompletionTokens,
		TotalTokens:      eventData.TotalTokens,
		MessagePayload:   eventData.MessagePayload,
		InputCost:        eventData.InputCost,
		OutputCost:       eventData.OutputCost,
		TotalCost:        eventData.TotalCost,
		CreatedAt:        eventData.CreatedAt,
		UpdatedAt:        eventData.UpdatedAt,
	}
}
