package command

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/engine/handler"
	"github.com/manboster/manboster/internal/engine/onboard"
	"github.com/manboster/manboster/internal/engine/safeguard"
	"github.com/manboster/manboster/internal/engine/soul"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/session"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

type required interface {
	BuildMessageRunner(instance chat.Provider, sessionId string)
}

type Handler struct {
	repo             repository.Repository
	onboard          *onboard.Service
	safeguardService *safeguard.Service
	sessionService   *session.Service
	llmProviders     map[string]llm.Provider
	config           *config.Config
	soulService      *soul.Service
	handler          *handler.Handler
	provider         *Provider[chat.CommandType]
	engine           required
}

func NewHandler(engine required, repo repository.Repository, safeguard *safeguard.Service, sessionService *session.Service, llmProviders map[string]llm.Provider, config *config.Config, soul *soul.Service, onboardService *onboard.Service, handler *handler.Handler) *Handler {
	h := &Handler{
		repo:             repo,
		safeguardService: safeguard,
		sessionService:   sessionService,
		llmProviders:     llmProviders,
		config:           config,
		soulService:      soul,
		onboard:          onboardService,
		handler:          handler,
		provider:         NewProvider[chat.CommandType](),
		engine:           engine,
	}
	h.Init()
	return h
}
