package types

import "time"

type Cron struct {
	ID           uint64    `gorm:"primary_key;auto_increment;column:id"`
	Name         string    `gorm:"unique;column:name"`
	ChatID       string    `gorm:"column:chat_id"`
	ChatProvider string    `gorm:"column:chat_provider"`
	CronTab      string    `gorm:"column:cron_tab"`
	Prompt       string    `gorm:"column:prompt"`
	Type         string    `gorm:"column:type"`
	CreatedBy    string    `gorm:"column:created_by"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
