package types

import "time"

type Memory struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;column:id"`
	Key       string    `gorm:"uniqueIndex;column:key"`
	Value     string    `gorm:"column:value"`
	Scope     string    `gorm:"column:scope"` // json array, listing user names
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
