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

func (repo *CronRepo) GetCronjobByName(ctx context.Context, name string) (types.Cron, error) {
	var cronDatabase dbtypes.Cron
	resp := repo.db.WithContext(ctx).Where("name = ?", name).First(&cronDatabase)
	if resp.Error != nil {
		return types.Cron{}, resp.Error
	}
	return types.MapCron(cronDatabase), nil
}

func (repo *CronRepo) UpdateCronjobTab(ctx context.Context, name string, newTab string) error {
	var cronDatabase dbtypes.Cron
	resp := repo.db.WithContext(ctx).Model(&cronDatabase).Where("name = ?", name).Update("cron_tab", newTab)
	if resp.RowsAffected == 0 {
		return ErrNotFound
	}
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

func (repo *CronRepo) UpdateCronjobPrompt(ctx context.Context, name string, newPrompt string) error {
	var cronDatabase dbtypes.Cron
	resp := repo.db.WithContext(ctx).Model(&cronDatabase).Where("name = ?", name).Update("prompt", newPrompt)
	if resp.RowsAffected == 0 {
		return ErrNotFound
	}
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

func (repo *CronRepo) DeleteCronjob(ctx context.Context, name string) error {
	resp := repo.db.WithContext(ctx).Where("name = ?", name).Delete(&dbtypes.Cron{})
	if resp.RowsAffected == 0 {
		return ErrNotFound
	}
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}
