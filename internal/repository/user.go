package repository

import (
	"context"

	dbtypes "github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/internal/repository/types"
	"gorm.io/gorm"
)

type UserRepository interface {
	UserCounts(ctx context.Context) (int64, error) // get user's counts
	UserInfo(ctx context.Context, platform string, id string) (types.User, error)
	CreateUser(ctx context.Context, user types.User) error
	DeleteUser(ctx context.Context, platform string, id string) error
}

type UserRepo struct {
	db *gorm.DB
}

// UserCounts gets the total number of users.
func (repo *UserRepo) UserCounts(ctx context.Context) (int64, error) {
	var count int64

	err := repo.db.WithContext(ctx).Model(&dbtypes.User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// UserInfo returns the info of user via userid
func (repo *UserRepo) UserInfo(ctx context.Context, platform string, id string) (types.User, error) {
	var user dbtypes.User

	err := repo.db.WithContext(ctx).Where("userid = ? AND platform = ?", id, platform).First(&user).Error
	if err != nil {
		return types.User{}, err
	}
	return types.MapUser(user), nil
}

// CreateUser adds user
func (repo *UserRepo) CreateUser(ctx context.Context, user types.User) error {
	uInfo := types.MapU(user)
	return repo.db.WithContext(ctx).Create(&uInfo).Error
}

// DeleteUser deletes authenticated user from this repository
func (repo *UserRepo) DeleteUser(ctx context.Context, platform string, id string) error {
	var uData dbtypes.User
	return repo.db.WithContext(ctx).Where("userid = ? AND platform = ?", id, platform).Delete(&uData).Error
}
