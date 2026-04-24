package types

import "time"

// ChatData <=> like llm.Message & llm.Model
type ChatData struct {
	ID               uint64    `gorm:"primary_key;auto_increment;column:id"`
	SessionID        string    `gorm:"column:session_id;index;not null"`
	Role             string    `gorm:"column:role"`
	MessageType      int16     `gorm:"column:message_type"`
	PromptTokens     int       `gorm:"column:prompt_tokens;default:0"`
	CompletionTokens int       `gorm:"column:completion_tokens;default:0"`
	TotalTokens      int       `gorm:"column:total_tokens;default:0"`
	InputCost        float64   `gorm:"column:input_cost;default:0"`
	OutputCost       float64   `gorm:"column:output_cost;default:0"`
	TotalCost        float64   `gorm:"column:total_cost;default:0"`
	MessagePayload   string    `gorm:"column:message_payload;type:text"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
}
