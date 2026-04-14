package types

import "github.com/manboster/manboster/internal/database/types"

type UserType int16

const (
	UserUnknown UserType = iota
	UserAdmin   UserType = iota
	UserRoot    UserType = iota
)

type User struct {
	ID       uint64
	UserID   string
	Type     UserType
	Platform string
}

func MapU(u User) types.User {
	return types.User{
		ID:       u.ID,
		UserID:   u.UserID,
		Type:     int16(u.Type),
		Platform: u.Platform,
	}
}

func MapUser(u types.User) User {
	return User{
		ID:       u.ID,
		UserID:   u.UserID,
		Type:     UserType(u.Type),
		Platform: u.Platform,
	}
}
