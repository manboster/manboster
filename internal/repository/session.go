package repository

import (
	"context"

	dbtypes "github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/internal/repository/types"
	"gorm.io/gorm"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, session types.Session) error
	GetSession(ctx context.Context, sessionId string) (types.Session, error)
	GetSessions(ctx context.Context) ([]types.Session, error)
	UpdateSession(ctx context.Context, sid string, updates map[string]interface{}) error
	DeleteSession(ctx context.Context, sessionId string) error
	GetAllSessions(ctx context.Context) ([]types.Session, error)
}

type SessionRepo struct {
	db *gorm.DB
}

// CreateSession creates session for a chat.
func (repo *SessionRepo) CreateSession(ctx context.Context, session types.Session) error {
	sessDBType := types.MapSess(session)
	return repo.db.WithContext(ctx).Create(&sessDBType).Error
}

// GetSession gets session data
func (repo *SessionRepo) GetSession(ctx context.Context, sessionId string) (types.Session, error) {
	var sessDBType dbtypes.Session
	err := repo.db.WithContext(ctx).Where("session_id = ?", sessionId).First(&sessDBType).Error
	if err != nil {
		return types.Session{}, err
	}
	return types.MapSession(sessDBType), nil
}

// GetSessions return first 20 session data
func (repo *SessionRepo) GetSessions(ctx context.Context) ([]types.Session, error) {
	var dbSessions []dbtypes.Session
	var s []types.Session
	resp := repo.db.WithContext(ctx).Order("created_at DESC").Find(&dbSessions).Limit(20)
	if resp.Error != nil {
		return nil, resp.Error
	}
	for _, session := range dbSessions {
		s = append(s, types.MapSession(session))
	}
	return s, nil
}

// GetAllSessions returns all sessions data
func (repo *SessionRepo) GetAllSessions(ctx context.Context) ([]types.Session, error) {
	var dbSessions []dbtypes.Session
	var s []types.Session
	resp := repo.db.WithContext(ctx).Order("created_at DESC").Find(&dbSessions)
	if resp.Error != nil {
		return nil, resp.Error
	}
	for _, session := range dbSessions {
		s = append(s, types.MapSession(session))
	}
	return s, nil
}

// UpdateSession updates session data
func (repo *SessionRepo) UpdateSession(ctx context.Context, sid string, updates map[string]interface{}) error {
	var count int64
	if err := repo.db.Model(&dbtypes.Session{}).Where("session_id = ?", sid).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return ErrNotFound
	}

	resp := repo.db.WithContext(ctx).Model(&dbtypes.Session{}).Where("session_id = ?", sid).Updates(updates)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

// DeleteSession deletes session data
func (repo *SessionRepo) DeleteSession(ctx context.Context, sessionId string) error {
	return repo.db.WithContext(ctx).Where("session_id = ?", sessionId).Delete(&dbtypes.Session{}).Error
}
