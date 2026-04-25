package types

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/database/types"
)

type Soul struct {
	ID        uint64
	Priority  uint8
	Name      string
	UserID    string
	Provider  string
	Scope     []string
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string
}

func MapS(soul Soul) types.Soul {
	val, err := json.Marshal(soul.Scope)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Repository] We encountered an error when encoing to grom object soul: %v, Object info: %+v", err, soul))
		return types.Soul{}
	}
	return types.Soul{
		ID:        soul.ID,
		Priority:  soul.Priority,
		Name:      soul.Name,
		UserID:    soul.UserID,
		Provider:  soul.Provider,
		Scope:     string(val),
		CreatedAt: soul.CreatedAt,
		UpdatedAt: soul.UpdatedAt,
		Content:   soul.Content,
	}
}

func MapSoul(soul types.Soul) Soul {
	var sl Soul
	err := json.Unmarshal([]byte(soul.Scope), &sl.Scope)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Repository] We encountered an error when encoing to grom object soul: %v, Object info: %+v", err, soul))
		return Soul{}
	}
	sl.ID = soul.ID
	sl.Priority = soul.Priority
	sl.Name = soul.Name
	sl.UserID = soul.UserID
	sl.Provider = soul.Provider
	sl.CreatedAt = soul.CreatedAt
	sl.UpdatedAt = soul.UpdatedAt
	sl.Content = soul.Content
	return sl
}
