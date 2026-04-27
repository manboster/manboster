package ignorance

import "sync"

type Manager struct {
	ignoreMark map[string]bool
	lock       *sync.RWMutex
}

func New() *Manager {
	return &Manager{
		ignoreMark: make(map[string]bool),
		lock:       &sync.RWMutex{},
	}
}
