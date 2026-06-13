package selection

import (
	"sync"

	"github.com/manboster/manboster/spec/chat"
)

type Manager struct {
	SelectionLock *sync.Mutex
	Selection     map[string]*chat.Message
	SelectionChan map[string]chan *chat.Message
}

func New() *Manager {
	return &Manager{
		SelectionLock: &sync.Mutex{},
		Selection:     make(map[string]*chat.Message),
		SelectionChan: make(map[string]chan *chat.Message),
	}
}
