package types

import "time"

type Soul struct {
	ID       uint64    `gorm:"primary_key;auto_increment;column:id"`
	Priority uint8     `gorm:"column:priority"`
	UserID   string    `gorm:"column:user_id"` // who injected this to system prompt?
	Scope    string    `gorm:"column:scope"`
	Time     time.Time `gorm:"column:time"`    // the Unix time when this injected to system prompt.
	Content  string    `gorm:"column:content"` // injected content
}
