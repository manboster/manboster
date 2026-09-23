package repository

import (
	"context"

	dbtypes "github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/spec/llm"
	"gorm.io/gorm"
)

type EventDataRepository interface {
	CreateEventData(ctx context.Context, chatData types.EventData) error
	GetEventData(ctx context.Context, sessionId string) ([]types.EventData, error)
	DeleteEventData(ctx context.Context, sessionId string) error
	CountEventDataTokenBySession(ctx context.Context, sessionId string) (llm.Usage, error)
	GetTotalToken(ctx context.Context, sessionId string) (int, error)
}

type EventDataRepo struct {
	db *gorm.DB
}

// CreateEventData creates chats data
func (repo *EventDataRepo) CreateEventData(ctx context.Context, chatData types.EventData) error {
	dbEventDataType := types.MapCD(chatData)
	return repo.db.WithContext(ctx).Create(&dbEventDataType).Error
}

// GetEventData gets chats' data from database
func (repo *EventDataRepo) GetEventData(ctx context.Context, sessionId string) ([]types.EventData, error) {
	var dbEventData []dbtypes.EventData
	var chatData []types.EventData

	// get chat data list from session ids
	resp := repo.db.WithContext(ctx).Where("session_id = ?", sessionId).Find(&dbEventData)
	if resp.Error != nil {
		return nil, resp.Error
	}

	// iterate to build a raw chat data array
	for _, dbEventDataVal := range dbEventData {
		chatData = append(chatData, types.MapEventData(dbEventDataVal))
	}
	return chatData, nil
}

// DeleteEventData deletes chats data via sessionId
func (repo *EventDataRepo) DeleteEventData(ctx context.Context, sessionId string) error {
	return repo.db.WithContext(ctx).Where("session_id = ?", sessionId).Delete(&dbtypes.EventData{}).Error
}

// CountEventDataTokenBySession counts all input/output tokens used in this chat session
func (repo *EventDataRepo) CountEventDataTokenBySession(ctx context.Context, sessionId string) (llm.Usage, error) {
	data, err := repo.GetEventData(ctx, sessionId)
	if err != nil {
		return llm.Usage{}, err
	}

	var usage llm.Usage
	usage.TotalTokens = 0
	usage.CompletionTokens = 0
	usage.PromptTokens = 0
	usage.InputCost = 0
	usage.OutputCost = 0
	usage.TotalCost = 0

	for _, dbEventData := range data {
		usage.PromptTokens += dbEventData.PromptTokens
		usage.CompletionTokens += dbEventData.CompletionTokens
		usage.TotalTokens += dbEventData.TotalTokens
		usage.InputCost += dbEventData.InputCost
		usage.OutputCost += dbEventData.OutputCost
		usage.TotalCost += dbEventData.TotalCost
	}
	return usage, nil
}

// GetTotalToken gets latest token used in this chat session
func (repo *EventDataRepo) GetTotalToken(ctx context.Context, sessionId string) (int, error) {
	var dbEventData []dbtypes.EventData
	resp := repo.db.WithContext(ctx).Where("session_id = ?", sessionId).Order("created_at DESC").Find(&dbEventData)
	if resp.Error != nil {
		return -1, resp.Error
	}

	for _, chatData := range dbEventData {
		if chatData.TotalTokens > 0 {
			return chatData.TotalTokens, nil
		}
	}
	return 0, nil
}
