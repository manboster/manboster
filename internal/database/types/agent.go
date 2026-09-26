package types

import "time"

// Agent stores abstract agents in the sqlite database, in order to enhance harness buildup
type Agent struct {
	ID        uint64    `gorm:"primary_key;auto_increment;column:id"`
	AgentID   string    `gorm:"column:agent_id"`
	Data      string    `gorm:"column:data"`
	Father    string    `gorm:"column:father"`   // getting this agent's father
	Children  string    `gorm:"column:children"` // listing children's ids
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
