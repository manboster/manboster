package types

import (
	"time"

	"github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/spec/schema"
)

type User struct {
	ID        uint64
	UserID    string
	Type      schema.UserType
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
		Type:      schema.UserType(u.Type),
		Platform:  u.Platform,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
