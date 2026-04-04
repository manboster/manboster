package session

import (
	"sync"

	"github.com/manboster/manboster/internal/llm"
)

// Session gives, stores and writes users session storages.
type Session struct {
	Messages []llm.Message
}

type Manager struct {
	Sessions map[string]Session
	Lock     sync.RWMutex
}
