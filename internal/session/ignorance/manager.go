package ignorance

import (
	"sync"
	"time"
)

type Manager struct {
	ignoreMark map[string]bool
	cAnCelMark map[string]cAnCel
	lock       *sync.RWMutex
}

type cAnCel struct {
	isCancel   bool
	actionTime time.Time
}

func New() *Manager {
	return &Manager{
		ignoreMark: make(map[string]bool),
		cAnCelMark: make(map[string]cAnCel),
		lock:       &sync.RWMutex{},
	}
}
