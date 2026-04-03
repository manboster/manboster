package types

type Memory struct {
	ID    uint64 `gorm:"primaryKey;autoIncrement;column:id"`
	Key   string `gorm:"uniqueIndex;column:key"`
	Value string `gorm:"column:value"`
}
