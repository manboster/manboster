package gguf

import (
	"sync"

	"github.com/hybridgroup/yzma/pkg/llama"
)

type Manager struct {
	avail      bool
	availModel bool
	load       bool
	modelCtx   llama.Context
	model      llama.Model
	lock       sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		avail:      false,
		availModel: false,
		lock:       sync.RWMutex{},
		load:       false,
	}
}

func (m *Manager) Avail() bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.avail
}

func (m *Manager) AvailModel() bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.availModel
}

func (m *Manager) SetAvail(avail bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.avail = avail
}

func (m *Manager) SetAvailModel(avail bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.availModel = avail
}

func (m *Manager) IsReady() bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.avail && m.availModel
}

func (m *Manager) SetModelCtx(modelCtx llama.Context) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.modelCtx = modelCtx
}

func (m *Manager) SetModel(model llama.Model) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.model = model
}

func (m *Manager) ModelCtx() llama.Context {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.modelCtx
}

func (m *Manager) Model() llama.Model {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.model
}

func (m *Manager) SetLoad(load bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.load = load
}

func (m *Manager) Load() bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.load
}
