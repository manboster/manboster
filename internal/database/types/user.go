package types

type User struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement;column:id"`
	UserID   string `gorm:"column:userid"`   // user's id, regardless of names.
	Type     int16  `gorm:"column:usertype"` // user's type
	Platform string `gorm:"column:platform"` // provider platform's id
}
