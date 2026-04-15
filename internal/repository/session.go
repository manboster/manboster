package repository

import (
	"context"

	"github.com/manboster/manboster/internal/repository/types"
)

type SessionRepository interface {
}

func (repo *Repo) CreateSession(ctx context.Context, session types.Session) error {
	sessDBType := types.MapSess(session)
	return repo.db.Create(&sessDBType).Error
}

func (repo *Repo) GetSession(ctx context.Context, sessionId string) (types.Session, error) {
	return types.Session{}, nil
}

func (repo *Repo) GetAllSessions(ctx context.Context) ([]types.Session, error) {
	return []types.Session{}, nil
}

func (repo *Repo) UpdateSession(ctx context.Context, session types.Session) error {
	return nil
}

func (repo *Repo) DeleteSession(ctx context.Context, sessionId string) error {
	return nil
}
