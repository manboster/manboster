package types

import "time"

type Soul struct {
	ID        uint64    `gorm:"primary_key;auto_increment;column:id"`
	Priority  uint8     `gorm:"column:priority"`
	UserID    string    `gorm:"column:user_id"` // who injected this to system prompt?
	Scope     string    `gorm:"column:scope"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Content   string    `gorm:"column:content"` // injected content
}
