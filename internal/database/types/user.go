package types

type User struct {
	ID       uint64 `gorm:"primary_key;auto_increment;column:id"`
	UserID   string `gorm:"column:userid"`   // user's id, regardless of names.
	Type     int16  `gorm:"column:usertype"` // user's type
	Platform string `gorm:"column:platform"` // provider platform's id
}
