package repository

import (
	"context"
	"encoding/json"

	dbtypes "github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/internal/repository/types"
)

type MemoryRepository interface {
}

func (repo *Repo) CreateMemory(ctx context.Context, m types.Memory) error {
	dbMemory := types.MapMem(m)
	return repo.db.WithContext(ctx).Create(&dbMemory).Error
}

func (repo *Repo) GetMemory(ctx context.Context, key string) (types.Memory, error) {
	var m dbtypes.Memory
	err := repo.db.WithContext(ctx).Where("key = ?", key).First(&m).Error
	if err != nil {
		return types.Memory{}, err
	}
	return types.MapMemory(m), nil
}

func (repo *Repo) EditMemoryScope(ctx context.Context, key string, scope []string) error {
	jsonify, err := json.Marshal(scope)
	if err != nil {
		return err
	}
	return repo.db.WithContext(ctx).Model(&dbtypes.Memory{}).Where("key = ?", key).Update("scope", string(jsonify)).Error
}

func (repo *Repo) AppendMemoryScope(ctx context.Context, key string, scope string) error {
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

func (repo *Repo) DeleteMemoryScope(ctx context.Context, key string, scope string) error {
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

func (repo *Repo) EditMemoryValue(ctx context.Context, key string, value string) error {
	return repo.db.WithContext(ctx).Model(&dbtypes.Memory{}).Where("key = ?", key).Update("value", value).Error
}

func (repo *Repo) DeleteMemory(ctx context.Context, key string) error {
	return repo.db.WithContext(ctx).Where("key = ?", key).Delete(&dbtypes.Memory{}).Error
}
