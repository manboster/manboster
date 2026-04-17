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
	config         *config.Config
	repo           repository.Repository

	onboardLock sync.Mutex
	pairKey     int64
	retry       int
	userCount   int64
}

func New(cfg *config.Config, repo repository.Repository, llmProviders []llm.Provider) (*Engine, error) {
	return &Engine{
		sessionManager: session.NewManager(),
		llmProviders:   llmProviders,
		config:         cfg,
		repo:           repo,
		pairKey:        0,
		retry:          0,
		userCount:      0,
	}, nil
}
