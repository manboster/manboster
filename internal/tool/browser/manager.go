package browser

import (
	"sync"

	"github.com/go-rod/rod"
)

type Manager struct {
	lock             sync.Mutex
	browserInstances map[string]*rod.Browser
}

func NewManager() *Manager {
	return &Manager{
		browserInstances: make(map[string]*rod.Browser),
		lock:             sync.Mutex{},
	}
}
