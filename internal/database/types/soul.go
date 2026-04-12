package types

type Soul struct {
	ID       uint64 `gorm:"primary_key;auto_increment;column:id"`
	Priority uint8  `gorm:"column:priority"`
	UserID   string `gorm:"column:user_id"`  // who injected this to system prompt?
	Platform string `gorm:"column:platform"` // which platform?
	ChatID   string `gorm:"column:chat_id"`  // which chats would be available?
	Time     int64  `gorm:"column:time"`     // the Unix time when this injected to system prompt.
	Content  string `gorm:"column:content"`  // injected content
}
