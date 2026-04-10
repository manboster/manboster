package engine

import (
	"sync"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/session"
)

type Engine struct {
	sessionManager *session.Manager
	llmProviders   []llm.Provider
	config         config.Config
	repo           repository.Repository

	onboardLock sync.Mutex
	pairKey     int64
	retry       int
}

func New(cfg config.Config, repo repository.Repository) (*Engine, error) {
	return &Engine{
		sessionManager: session.NewManager(),
		llmProviders:   []llm.Provider{},
		config:         cfg,
		repo:           repo,
		pairKey:        0,
		retry:          0,
	}, nil
}
