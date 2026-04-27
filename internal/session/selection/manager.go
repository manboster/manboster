package selection

import (
	"sync"

	"github.com/manboster/manboster/spec/chat"
)

type Manager struct {
	SelectionLocks map[string]*sync.Mutex
	Selection      map[string]*chat.Message
}

func New() *Manager {
	return &Manager{
		SelectionLocks: make(map[string]*sync.Mutex),
		Selection:      make(map[string]*chat.Message),
	}
}
