package session

import (
	"context"
	"sync"

	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/llm"
)

// Session gives, stores and writes users session storages.
type Session struct {
	Events      []llm.Event
	Provider    string
	Model       string
	Command     chat.CommandType // if command is triggered, it would not be empty
	CommandStep int8             // the current step command is executing
	Active      bool
	Cancel      context.CancelFunc
}

type Manager struct {
	Sessions         map[string]Session
	Lock             sync.RWMutex
	SessionLocks     map[string]*sync.Mutex
	SessionChatLocks map[string]*sync.Mutex
}
