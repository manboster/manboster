package engine

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/session"
)

type Engine struct {
	sessionManager *session.Manager
	llmProviders   []llm.Provider
	config         config.Config
}

func New(cfg config.Config) (*Engine, error) {
	return &Engine{
		sessionManager: session.NewManager(),
		llmProviders:   []llm.Provider{},
		config:         cfg,
	}, nil
}
