package repository

import (
	"context"

	"github.com/manboster/manboster/internal/database/types"
)

type UserRepository interface {
	UserCounts(ctx context.Context) (int64, error) // get user's counts
}

// UserCounts gets the total number of users.
func (repo *Repo) UserCounts(ctx context.Context) (int64, error) {
	var count int64

	err := repo.db.Model(&types.User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
