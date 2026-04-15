package repository

import (
	"context"

	dbtypes "github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/internal/repository/types"
)

type ChatRepository interface {
	CreateChat(chat types.Chat) error
	GetChat(ctx context.Context, chatId string, provider string) (types.Chat, error)
	GetAllChats(ctx context.Context) ([]types.Chat, error)
	DeleteChat(ctx context.Context, chatId string, provider string) error
}

// CreateChat creates a new chat information
func (repo *Repo) CreateChat(ctx context.Context, chat types.Chat) error {
	dbChat := types.MapC(chat)
	return repo.db.WithContext(ctx).Create(&dbChat).Error
}

// GetChat gets chat information via chatId and provider
func (repo *Repo) GetChat(ctx context.Context, chatId string, provider string) (types.Chat, error) {
	var dbChatInfo dbtypes.Chat
	err := repo.db.WithContext(ctx).Where("chat_id = ? AND chat_provider = ?", chatId, provider).First(&dbChatInfo).Error
	if err != nil {
		return types.Chat{}, err
	}
	return types.MapChat(dbChatInfo), nil
}

// GetAllChats TODO: gets all chat's information
func (repo *Repo) GetAllChats(ctx context.Context) ([]types.Chat, error) {
	//var dbChatInfo []dbtypes.Chat
	//err := repo.db.Model(&dbChatInfo).Error
	//if err != nil {
	//	return []types.Chat{}, err
	//}
	//return types.MapChat(dbChatInfo), nil
	return []types.Chat{}, nil
}

// DeleteChat deletes chat information by chatId and Provider
func (repo *Repo) DeleteChat(ctx context.Context, chatId string, provider string) error {
	return repo.db.WithContext(ctx).Where("chat_id = ? AND chat_provider = ?", chatId, provider).Delete(&dbtypes.Chat{}).Error
}
