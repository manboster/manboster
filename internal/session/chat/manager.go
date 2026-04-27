package chat

import "sync"

type Manager struct {
	SessionChatLocks map[string]*sync.Mutex
}

func New() *Manager {
	return &Manager{
		SessionChatLocks: make(map[string]*sync.Mutex),
	}
}
