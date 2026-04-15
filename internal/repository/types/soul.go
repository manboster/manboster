package types

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/database/types"
)

type Soul struct {
	ID       uint64
	Priority uint8
	UserID   string
	Scope    []string
	Time     time.Time
	Content  string
}

func MapS(soul Soul) types.Soul {
	val, err := json.Marshal(soul.Scope)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Repository] We encountered an error when encoing to grom object soul: %v, Object info: %+v", err, soul))
		return types.Soul{}
	}
	return types.Soul{
		ID:       soul.ID,
		Priority: soul.Priority,
		UserID:   soul.UserID,
		Scope:    string(val),
		Time:     soul.Time,
		Content:  soul.Content,
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
	sl.UserID = soul.UserID
	sl.Time = soul.Time
	sl.Content = soul.Content
	return sl
}
