package engine

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/engine/onboard"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/session"
)

type Engine struct {
	sessionManager *session.Manager
	llmProviders   []llm.Provider
	config         *config.Config
	repo           repository.Repository
	onboard        *onboard.Service
}

func New(cfg *config.Config, repo repository.Repository, llmProviders []llm.Provider) (*Engine, error) {
	return &Engine{
		sessionManager: session.NewManager(),
		llmProviders:   llmProviders,
		config:         cfg,
		repo:           repo,
		onboard:        nil,
	}, nil
}
