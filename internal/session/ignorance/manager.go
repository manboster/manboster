package ignorance

import (
	"sync"
	"time"
)

type Manager struct {
	ignoreMark map[string]mark
	cAnCelMark map[string]mark
	lock       *sync.RWMutex
}

type mark struct {
	m          bool
	actionTime time.Time
	ttl        int
}

func New() *Manager {
	return &Manager{
		ignoreMark: make(map[string]mark),
		cAnCelMark: make(map[string]mark),
		lock:       &sync.RWMutex{},
	}
}
