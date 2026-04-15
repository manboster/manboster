package types

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/database/types"
)

type Memory struct {
	ID        uint64
	Key       string
	Value     string
	Scope     []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func MapMem(memory Memory) types.Memory {
	val, err := json.Marshal(memory.Scope)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Repository] We encountered an error when converting type memory to gorm object: %v; Memory Object data: %+v", err, memory))
		return types.Memory{}
	}
	return types.Memory{
		ID:        memory.ID,
		Key:       memory.Key,
		Value:     memory.Value,
		Scope:     string(val),
		CreatedAt: memory.CreatedAt,
		UpdatedAt: memory.UpdatedAt,
	}
}

func MapMemory(memory types.Memory) Memory {
	var mem Memory
	err := json.Unmarshal([]byte(memory.Scope), &mem.Scope)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Repository] We encountered an error when converting gorm object to memory type: %v; Gorm Object data: %+v", err, memory))
	}
	mem.ID = memory.ID
	mem.Key = memory.Key
	mem.Value = memory.Value
	mem.CreatedAt = memory.CreatedAt
	mem.UpdatedAt = memory.UpdatedAt
	return mem
}
