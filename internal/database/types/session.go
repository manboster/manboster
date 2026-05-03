package types

import "time"

type Session struct {
	ID               uint64    `gorm:"primaryKey;autoIncrement;column:id"` // session's auto-increment ids
	SessionID        string    `gorm:"uniqueIndex;column:session_id"`      // identifiable session id (maybe 16 keys?)
	LLMProviderModel string    `gorm:"column:llm_provider_model"`          //
	LLMProvider      string    `gorm:"column:llm_provider"`
	ActivatedSouls   string    `gorm:"column:activated_souls"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
}
