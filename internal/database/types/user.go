package types

import "time"

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;column:id"`
	UserID    string    `gorm:"column:userid"`   // user's id, regardless of names.
	Type      int16     `gorm:"column:usertype"` // user's type
	Platform  string    `gorm:"column:platform"` // provider platform's id
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
