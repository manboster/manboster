package types

type Chat struct {
	ID                uint64 `gorm:"primary_key;auto_increment;column:id"`
	ChatID            string `gorm:"column:chat_id"`
	ChatProviderID    uint64 `gorm:"column:chat_provider_id"`
	ChatProviderModel string `gorm:"column:chat_provider_model"`
}
