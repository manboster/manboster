package repository

import (
	"context"

	dbtypes "github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/internal/repository/types"
	"gorm.io/gorm"
)

type CronRepository interface {
}

type CronRepo struct {
	db *gorm.DB
}

func NewCronRepository(db *gorm.DB) CronRepository {
	return &CronRepo{
		db: db,
	}
}

func (repo *CronRepo) CreateCronjob(ctx context.Context, cj types.Cron) error {
	cronjobDatabaseType := types.MapCr(cj)
	return repo.db.WithContext(ctx).Create(&cronjobDatabaseType).Error
}

func (repo *CronRepo) GetCronjobByChatID(ctx context.Context, chat string, provider string) ([]types.Cron, error) {
	var cronjobDB []dbtypes.Cron
	var cj []types.Cron
	resp := repo.db.WithContext(ctx).Where("chat_id = ? AND chat_provider = ?", chat, provider).Find(&cronjobDB)
	if resp.Error != nil {
		return nil, resp.Error
	}
	for _, cronjob := range cronjobDB {
		cj = append(cj, types.MapCron(cronjob))
	}
	return cj, nil
}
