package chat

import "sync"

type Manager struct {
	SessionChatLocks map[string]*sync.Mutex

	SessionChatData map[string]data
	Lock            sync.RWMutex
}

type data struct {
	Counter int
	MsgId   string
}

func New() *Manager {
	return &Manager{
		SessionChatLocks: make(map[string]*sync.Mutex),
	}
}
