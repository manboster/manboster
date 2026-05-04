package repository

import (
	"context"

	dbtypes "github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/spec/llm"
	"gorm.io/gorm"
)

type ChatDataRepository interface {
	CreateChatData(ctx context.Context, chatData types.ChatData) error
	GetChatData(ctx context.Context, sessionId string) ([]types.ChatData, error)
	DeleteChatData(ctx context.Context, sessionId string) error
	CountChatDataTokenBySession(ctx context.Context, sessionId string) (llm.Usage, error)
	GetTotalToken(ctx context.Context, sessionId string) (int, error)
}

type ChatDataRepo struct {
	db *gorm.DB
}

// CreateChatData creates chats data
func (repo *ChatDataRepo) CreateChatData(ctx context.Context, chatData types.ChatData) error {
	dbChatDataType := types.MapCD(chatData)
	return repo.db.WithContext(ctx).Create(&dbChatDataType).Error
}

// GetChatData gets chats' data from database
func (repo *ChatDataRepo) GetChatData(ctx context.Context, sessionId string) ([]types.ChatData, error) {
	var dbChatData []dbtypes.ChatData
	var chatData []types.ChatData

	// get chat data list from session ids
	resp := repo.db.WithContext(ctx).Where("session_id = ?", sessionId).Find(&dbChatData)
	if resp.Error != nil {
		return nil, resp.Error
	}

	// iterate to build a raw chat data array
	for _, dbChatDataVal := range dbChatData {
		chatData = append(chatData, types.MapChatData(dbChatDataVal))
	}
	return chatData, nil
}

// DeleteChatData deletes chats data via sessionId
func (repo *ChatDataRepo) DeleteChatData(ctx context.Context, sessionId string) error {
	return repo.db.WithContext(ctx).Where("session_id = ?", sessionId).Delete(&dbtypes.ChatData{}).Error
}

// CountChatDataTokenBySession counts all input/output tokens used in this chat session
func (repo *ChatDataRepo) CountChatDataTokenBySession(ctx context.Context, sessionId string) (llm.Usage, error) {
	data, err := repo.GetChatData(ctx, sessionId)
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

	for _, dbChatData := range data {
		usage.PromptTokens += dbChatData.PromptTokens
		usage.CompletionTokens += dbChatData.CompletionTokens
		usage.TotalTokens += dbChatData.TotalTokens
		usage.InputCost += dbChatData.InputCost
		usage.OutputCost += dbChatData.OutputCost
		usage.TotalCost += dbChatData.TotalCost
	}
	return usage, nil
}

// GetTotalToken gets latest token used in this chat session
func (repo *ChatDataRepo) GetTotalToken(ctx context.Context, sessionId string) (int, error) {
	var dbChatData []dbtypes.ChatData
	resp := repo.db.WithContext(ctx).Where("session_id = ?", sessionId).Order("created_at DESC").Find(&dbChatData)
	if resp.Error != nil {
		return -1, resp.Error
	}

	for _, chatData := range dbChatData {
		if chatData.TotalTokens > 0 {
			return chatData.TotalTokens, nil
		}
	}
	return 0, nil
}
