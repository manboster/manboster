package skills

import (
	"context"
	"errors"

	"github.com/manboster/manboster/internal/database"
	"github.com/manboster/manboster/internal/repository"
)

var sessionDB repository.SessionRepository

func (s *Service) Init(ctx context.Context, conf any) error {
	if sessionDB != nil {
		return nil
	}

	if database.DBInstance == nil || database.DBInstance.Instance() == nil {
		return errors.New("database not initialized")
	}
	sessionDB = repository.New(database.DBInstance.Instance())
	return nil
}
