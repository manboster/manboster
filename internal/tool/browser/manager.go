package browser

import (
	"context"
	"fmt"
	"path/filepath"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/manboster/manboster/internal/config"
)

type Manager struct {
	lock             sync.Mutex
	browserInstances map[string]*Instance
}

type Instance struct {
	browser  *rod.Browser
	lastUsed time.Time
	cAnCel   context.CancelFunc
	active   bool
}

func NewManager() *Manager {
	return &Manager{
		browserInstances: make(map[string]*Instance),
		lock:             sync.Mutex{},
	}
}

func (m *Manager) getBrowserInstance(ctx context.Context, id string) (*Instance, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	i, avail := m.browserInstances[id]
	if avail {
		i.lastUsed = time.Now()
		m.browserInstances[id] = i
		return i, nil
	}
	i = &Instance{}
	l := launcher.New().
		UserDataDir(config.Path(filepath.Join("browser", fmt.Sprintf("session-%s", id)))).
		Headless(false).
		Devtools(true)

	url, err := l.Launch()
	if err != nil {
		l.Kill()
		return nil, err
	}

	cancelCtx, cancel := context.WithCancel(ctx)
	browser := rod.New().ControlURL(url).Context(cancelCtx)
	go func(la *launcher.Launcher, cAnCel context.CancelFunc) {
		err := m.timeCheckRunner(ctx, id)
		if err != nil {
			color.Yellow("[Manboster Tool Provider] We encountered an error when stopping the browser")
		}
		la.Kill()
		cAnCel()
	}(l, cancel)

	err = browser.Connect()
	if err != nil {
		l.Kill()
		return nil, err
	}

	i.browser = browser
	i.lastUsed = time.Now()
	i.active = true
	i.cAnCel = cancel
	m.browserInstances[id] = i
	return i, nil
}

func (m *Manager) timeCheckRunner(ctx context.Context, id string) error {
	ticker := time.NewTicker(1 * time.Minute)
	for {
		select {
		case <-ticker.C:
			m.lock.Lock()
			i, found := m.browserInstances[id]
			if found && i.active {
				if time.Since(i.lastUsed) > 10*time.Minute {
					err := i.browser.Close()
					if err != nil {
						return err
					}
					delete(m.browserInstances, id)
				}
			}
			m.lock.Unlock()
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
