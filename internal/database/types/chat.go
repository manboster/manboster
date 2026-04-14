package types

type Chat struct {
	ID             uint64 `gorm:"primary_key;auto_increment;column:id"`
	ChatID         string `gorm:"column:chat_id"`         // chat's id
	ChatProvider   uint64 `gorm:"column:chat_provider"`   // chat's provider
	ChatModel      string `gorm:"column:chat_model"`      // chat's model
	ChatPermission int16  `gorm:"column:chat_permission"` // chat's permission
	SessionID      string `gorm:"column:session_id"`      // chat's session id, n chat ids to session ids
}
