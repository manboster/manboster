package repository

import (
	"context"

	dbtypes "github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/repository/types"
)

type ChatDataRepository interface {
	CreateChatData(ctx context.Context, chatData types.ChatData) error
	GetChatData(ctx context.Context, sessionId string) ([]types.ChatData, error)
	DeleteChatData(ctx context.Context, sessionId string) error
	CountChatDataTokenBySession(ctx context.Context, sessionId string) (llm.Usage, error)
	GetTotalToken(ctx context.Context, sessionId string) (int, error)
}

// CreateChatData creates chats data
func (repo *Repo) CreateChatData(ctx context.Context, chatData types.ChatData) error {
	dbChatDataType := types.MapCD(chatData)
	return repo.db.WithContext(ctx).Create(&dbChatDataType).Error
}

// GetChatData gets chats' data from database
func (repo *Repo) GetChatData(ctx context.Context, sessionId string) ([]types.ChatData, error) {
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
func (repo *Repo) DeleteChatData(ctx context.Context, sessionId string) error {
	return repo.db.WithContext(ctx).Where("session_id = ?", sessionId).Delete(&dbtypes.ChatData{}).Error
}

// CountChatDataTokenBySession counts all input/output tokens used in this chat session
func (repo *Repo) CountChatDataTokenBySession(ctx context.Context, sessionId string) (llm.Usage, error) {
	data, err := repo.GetChatData(ctx, sessionId)
	if err != nil {
		return llm.Usage{}, err
	}

	var usage llm.Usage
	usage.TotalTokens = 0
	usage.CompletionTokens = 0
	usage.PromptTokens = 0

	for _, dbChatData := range data {
		usage.PromptTokens += dbChatData.PromptTokens
		usage.CompletionTokens += dbChatData.CompletionTokens
		usage.TotalTokens += dbChatData.TotalTokens
	}
	return usage, nil
}

// GetTotalToken gets latest token used in this chat session
func (repo *Repo) GetTotalToken(ctx context.Context, sessionId string) (int, error) {
	var dbChatData dbtypes.ChatData
	resp := repo.db.WithContext(ctx).Where("session_id = ?", sessionId).Order("created_at DESC").First(&dbChatData)
	if resp.Error != nil {
		return -1, resp.Error
	}
	return dbChatData.TotalTokens, nil
}
