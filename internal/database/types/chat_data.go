package types

// ChatData <=> like llm.Message & llm.Model
type ChatData struct {
	ID             uint64 `gorm:"primary_key;auto_increment;column:id"`
	SessionID      string `gorm:"column:session_id"`
	Role           string `gorm:"column:role"`
	MessageType    int16  `gorm:"column:message_type"`
	Tokens         int    `gorm:"column:tokens"`
	MessagePayload string `gorm:"column:message_payload"`
}
