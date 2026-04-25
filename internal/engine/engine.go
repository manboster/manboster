package engine

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/engine/chatdata"
	"github.com/manboster/manboster/internal/engine/onboard"
	"github.com/manboster/manboster/internal/engine/safeguard"
	"github.com/manboster/manboster/internal/engine/soul"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/session"
	"github.com/manboster/manboster/internal/tool"
)

type Engine struct {
	sessionManager *session.Manager
	llmProviders   map[string]llm.Provider
	toolProviders  []tool.Provider
	toolMaps       map[string]tool.Provider
	config         *config.Config
	repo           repository.Repository

	onboard          *onboard.Service
	safeguardService *safeguard.Service
	chatDataService  *chatdata.Service
	soulService      *soul.Service
}

func New(cfg *config.Config, repo repository.Repository, llmProviders map[string]llm.Provider, toolProviders []tool.Provider) (*Engine, error) {
	return &Engine{
		sessionManager: session.NewManager(),
		llmProviders:   llmProviders,
		toolProviders:  toolProviders,
		toolMaps:       make(map[string]tool.Provider),
		config:         cfg,
		repo:           repo,
		onboard:        nil,
	}, nil
}
