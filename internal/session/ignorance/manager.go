package ignorance

import (
	"sync"
)

type Manager struct {
	mark        map[string]mark
	hachimiMark map[string]bool
	lock        *sync.RWMutex
}

func New() *Manager {
	return &Manager{
		mark:        make(map[string]mark),
		lock:        &sync.RWMutex{},
		hachimiMark: make(map[string]bool),
	}
}
