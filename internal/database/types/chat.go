package types

type Chat struct {
	ID             uint64 `gorm:"primary_key;auto_increment;column:id"`
	ChatID         string `gorm:"column:chat_id"`
	ChatProviderID uint64 `gorm:"column:chat_provider_id"`
	ChatPermission int16  `gorm:"column:chat_permission"`
	SessionID      string `gorm:"column:session_id"`
}
