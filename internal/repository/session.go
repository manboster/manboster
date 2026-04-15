package repository

import (
	"context"

	dbtypes "github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/internal/repository/types"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, session types.Session) error
	GetSession(ctx context.Context, sessionId string) (types.Session, error)
	GetAllSessions(ctx context.Context) ([]types.Session, error)
	UpdateSession(ctx context.Context, session types.Session) error
	DeleteSession(ctx context.Context, sessionId string) error
}

// CreateSession creates session for a chat.
func (repo *Repo) CreateSession(ctx context.Context, session types.Session) error {
	sessDBType := types.MapSess(session)
	return repo.db.WithContext(ctx).Create(&sessDBType).Error
}

// GetSession gets session data
func (repo *Repo) GetSession(ctx context.Context, sessionId string) (types.Session, error) {
	var sessDBType dbtypes.Session
	err := repo.db.WithContext(ctx).Where("session_id = ?", sessionId).First(&sessDBType).Error
	if err != nil {
		return types.Session{}, err
	}
	return types.MapSession(sessDBType), nil
}

// GetAllSessions return all session data
func (repo *Repo) GetAllSessions(ctx context.Context) ([]types.Session, error) {
	return []types.Session{}, nil
}

// UpdateSession updates session data
func (repo *Repo) UpdateSession(ctx context.Context, sid string, updates map[string]interface{}) error {
	resp := repo.db.WithContext(ctx).Model(&dbtypes.Session{}).Where("session_id = ?", sid).Updates(updates)
	if resp.Error != nil {
		return resp.Error
	}
	if resp.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

// DeleteSession deletes session data
func (repo *Repo) DeleteSession(ctx context.Context, sessionId string) error {
	return repo.db.WithContext(ctx).Where("session_id = ?", sessionId).Delete(&types.Session{}).Error
}
