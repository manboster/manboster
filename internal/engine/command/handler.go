package command

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/engine/handler"
	"github.com/manboster/manboster/internal/engine/onboard"
	"github.com/manboster/manboster/internal/engine/safeguard"
	"github.com/manboster/manboster/internal/engine/soul"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/session"
	"github.com/manboster/manboster/spec/llm"
)

type Handler struct {
	repo             repository.Repository
	onboard          *onboard.Service
	safeguardService *safeguard.Service
	sessionManager   *session.Manager
	llmProviders     map[string]llm.Provider
	config           *config.Config
	soulService      *soul.Service
	handler          *handler.Handler
}

func NewHandler(repo repository.Repository, safeguard *safeguard.Service, sessionManager *session.Manager, llmProviders map[string]llm.Provider, config *config.Config, soul *soul.Service, onboardService *onboard.Service, handler *handler.Handler) *Handler {
	return &Handler{
		repo:             repo,
		safeguardService: safeguard,
		sessionManager:   sessionManager,
		llmProviders:     llmProviders,
		config:           config,
		soulService:      soul,
		onboard:          onboardService,
		handler:          handler,
	}
}

func (h *Handler) InjectOnboardSvc(onboard *onboard.Service) {
	h.onboard = onboard
}
