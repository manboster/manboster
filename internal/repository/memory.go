package repository

import (
	"context"
	"encoding/json"

	dbtypes "github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/internal/repository/types"
	"gorm.io/gorm"
)

type MemoryRepository interface {
	CreateMemory(ctx context.Context, m types.Memory) error
	GetMemory(ctx context.Context, key string) (types.Memory, error)
	EditMemoryScope(ctx context.Context, key string, scope []string) error
	AppendMemoryScope(ctx context.Context, key string, scope string) error
	DeleteMemoryScope(ctx context.Context, key string, scope string) error
	EditMemoryValue(ctx context.Context, key string, value string) error
	DeleteMemory(ctx context.Context, key string) error
	ListMemoryKeys(ctx context.Context) ([]string, error)
}

type MemoryRepo struct {
	db *gorm.DB
}

func NewMemoryRepo(db *gorm.DB) *MemoryRepo {
	return &MemoryRepo{
		db: db,
	}
}

func (repo *MemoryRepo) CreateMemory(ctx context.Context, m types.Memory) error {
	dbMemory := types.MapMem(m)
	return repo.db.WithContext(ctx).Create(&dbMemory).Error
}

func (repo *MemoryRepo) GetMemory(ctx context.Context, key string) (types.Memory, error) {
	var m dbtypes.Memory
	err := repo.db.WithContext(ctx).Where("key = ?", key).First(&m).Error
	if err != nil {
		return types.Memory{}, err
	}
	return types.MapMemory(m), nil
}

func (repo *MemoryRepo) EditMemoryScope(ctx context.Context, key string, scope []string) error {
	jsonify, err := json.Marshal(scope)
	if err != nil {
		return err
	}
	return repo.db.WithContext(ctx).Model(&dbtypes.Memory{}).Where("key = ?", key).Update("scope", string(jsonify)).Error
}

func (repo *MemoryRepo) AppendMemoryScope(ctx context.Context, key string, scope string) error {
	m, err := repo.GetMemory(ctx, key)
	if err != nil {
		return err
	}
	for _, v := range m.Scope {
		if v == scope {
			return ErrDuplicateMemoryScope
		}
	}
	m.Scope = append(m.Scope, scope)
	return repo.EditMemoryScope(ctx, key, m.Scope)
}

func (repo *MemoryRepo) DeleteMemoryScope(ctx context.Context, key string, scope string) error {
	m, err := repo.GetMemory(ctx, key)
	if err != nil {
		return err
	}
	for i, v := range m.Scope {
		if v == scope {
			m.Scope = append(m.Scope[:i], m.Scope[i+1:]...)
			return repo.EditMemoryScope(ctx, key, m.Scope)
		}
	}
	return ErrNotFound
}

func (repo *MemoryRepo) EditMemoryValue(ctx context.Context, key string, value string) error {
	return repo.db.WithContext(ctx).Model(&dbtypes.Memory{}).Where("key = ?", key).Update("value", value).Error
}

func (repo *MemoryRepo) DeleteMemory(ctx context.Context, key string) error {
	return repo.db.WithContext(ctx).Where("key = ?", key).Delete(&dbtypes.Memory{}).Error
}

func (repo *MemoryRepo) ListMemoryKeys(ctx context.Context) ([]string, error) {
	var m []dbtypes.Memory
	err := repo.db.WithContext(ctx).Find(&m).Error
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(m))
	for _, v := range m {
		keys = append(keys, v.Key)
	}
	return keys, nil
}
