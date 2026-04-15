package repository

import (
	"context"

	dbtypes "github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/internal/repository/types"
)

type ChatDataRepository interface {
	CreateChatData(ctx context.Context, chatData types.ChatData) error
	GetChatData(ctx context.Context, sessionId string) ([]types.ChatData, error)
	DeleteChatData(ctx context.Context, sessionId string) error
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
