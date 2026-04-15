package session

import (
	"context"
	"sync"

	"github.com/manboster/manboster/internal/llm"
)

// Session gives, stores and writes users session storages.
type Session struct {
	Events   []llm.Event
	Provider string
	Model    string
	Active   bool
	Cancel   context.CancelFunc
}

type Manager struct {
	Sessions         map[string]Session
	Lock             sync.RWMutex
	SessionLocks     map[string]*sync.Mutex
	SessionChatLocks map[string]*sync.Mutex
}
