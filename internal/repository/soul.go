package repository

import (
	"context"
	"encoding/json"

	dbtypes "github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/internal/repository/types"
)

type SoulRepository interface {
	CreateSoul(ctx context.Context, so types.Soul) error
	ReadSoulsByScope(ctx context.Context, scope string) ([]types.Soul, error)
	ReadSoulsByCreator(ctx context.Context, creator string) ([]types.Soul, error)
	GetSoul(ctx context.Context, name string) (types.Soul, error)
	GetAllSouls(ctx context.Context) ([]types.Soul, error)
	UpdateSoulScope(ctx context.Context, name string, scope []string) error
	UpdateSoulContent(ctx context.Context, name string, content string) error
	AppendSoulScope(ctx context.Context, name string, scope string) error
	DeleteSoul(ctx context.Context, name string) error
}

// CreateSoul helps you to create system prompt
func (repo *Repo) CreateSoul(ctx context.Context, so types.Soul) error {
	s := types.MapS(so)
	return repo.db.WithContext(ctx).Create(&s).Error
}

// ReadSoulsByScope helps you to read soul data by scope
func (repo *Repo) ReadSoulsByScope(ctx context.Context, scope string) ([]types.Soul, error) {
	var allSoul []dbtypes.Soul
	err := repo.db.WithContext(ctx).Find(&allSoul).Error
	if err != nil {
		return nil, err
	}

	var souls []types.Soul
	for _, soul := range allSoul {
		s := types.MapSoul(soul)
		for _, str := range s.Scope {
			if str == scope {
				souls = append(souls, s)
				break
			}
		}
	}

	return souls, nil
}

// ReadSoulsByCreator helps you to read soul data by creator
func (repo *Repo) ReadSoulsByCreator(ctx context.Context, creator string) ([]types.Soul, error) {
	var allSoul []dbtypes.Soul
	err := repo.db.WithContext(ctx).Where("user_id = ?", creator).Find(&allSoul).Error
	if err != nil {
		return nil, err
	}
	var souls []types.Soul
	for _, soul := range allSoul {
		souls = append(souls, types.MapSoul(soul))
	}
	return souls, nil
}

// GetSoul gets soul by its unique name
func (repo *Repo) GetSoul(ctx context.Context, name string) (types.Soul, error) {
	var soul dbtypes.Soul
	err := repo.db.WithContext(ctx).First(&soul, "name = ?", name).Error
	if err != nil {
		return types.Soul{}, err
	}
	return types.MapSoul(soul), nil
}

// GetAllSouls gets all souls data by getting its map
func (repo *Repo) GetAllSouls(ctx context.Context) ([]types.Soul, error) {
	var allSoul []dbtypes.Soul
	err := repo.db.WithContext(ctx).Find(&allSoul).Error
	if err != nil {
		return nil, err
	}
	var souls []types.Soul
	for _, soul := range allSoul {
		souls = append(souls, types.MapSoul(soul))
	}
	return souls, nil
}

// UpdateSoulScope helps you to update soul scope data of a soul.
func (repo *Repo) UpdateSoulScope(ctx context.Context, name string, scope []string) error {
	jsonify, _ := json.Marshal(scope)
	return repo.db.WithContext(ctx).Model(&dbtypes.Soul{}).Where("name = ?", name).Update("scope", string(jsonify)).Error
}

// UpdateSoulContent helps you to update soul content data of a soul.
func (repo *Repo) UpdateSoulContent(ctx context.Context, name string, content string) error {
	return repo.db.WithContext(ctx).Model(&dbtypes.Soul{}).Where("name = ?", name).Update("content", content).Error
}

// AppendSoulScope helps you to append soul scope by getting its id.
func (repo *Repo) AppendSoulScope(ctx context.Context, name string, scope string) error {
	s, err := repo.GetSoul(ctx, name)
	if err != nil {
		return err
	}
	for _, soul := range s.Scope {
		if soul == scope {
			return ErrDuplicateSoulScope
		}
	}
	s.Scope = append(s.Scope, scope)
	return repo.UpdateSoulScope(ctx, name, s.Scope)
}

// DeleteSoul deletes soul from the database
func (repo *Repo) DeleteSoul(ctx context.Context, name string) error {
	return repo.db.WithContext(ctx).Where("name = ?", name).Delete(&dbtypes.Soul{}).Error
}
