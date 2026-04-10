package repository

import (
	"context"

	dbtypes "github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/internal/repository/types"
)

type UserRepository interface {
	UserCounts(ctx context.Context) (int64, error) // get user's counts
	UserInfo(ctx context.Context, platform string, id string) (types.User, error)
	CreateUser(ctx context.Context, user types.User) error
}

// UserCounts gets the total number of users.
func (repo *Repo) UserCounts(ctx context.Context) (int64, error) {
	var count int64

	err := repo.db.Model(&dbtypes.User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// UserInfo returns the info of user via userid
func (repo *Repo) UserInfo(ctx context.Context, platform string, id string) (types.User, error) {
	var user dbtypes.User

	err := repo.db.Where("id = ? AND platform = ?", id, platform).First(&user).Error
	if err != nil {
		return types.User{}, err
	}
	return types.MapUser(user), nil
}

// CreateUser adds user
func (repo *Repo) CreateUser(ctx context.Context, user types.User) error {
	var uInfo dbtypes.User
	uInfo = types.MapU(user)
	return repo.db.Create(&uInfo).Error
}
