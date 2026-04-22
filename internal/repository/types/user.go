package types

import (
	"time"

	"github.com/manboster/manboster/internal/database/types"
)

type UserType int16

const (
	UserUnknown UserType = iota
	UserAdmin   UserType = iota
	UserRoot    UserType = iota
)

func (u UserType) String() string {
	switch u {
	case UserAdmin:
		return "admin"
	case UserRoot:
		return "root"
	default:
		return "unknown"
	}
}

type User struct {
	ID        uint64
	UserID    string
	Type      UserType
	Platform  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func MapU(u User) types.User {
	return types.User{
		ID:        u.ID,
		UserID:    u.UserID,
		Type:      int16(u.Type),
		Platform:  u.Platform,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func MapUser(u types.User) User {
	return User{
		ID:        u.ID,
		UserID:    u.UserID,
		Type:      UserType(u.Type),
		Platform:  u.Platform,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
