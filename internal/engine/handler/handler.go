package handler

import (
	"github.com/manboster/manboster/internal/engine/chatdata"
	"github.com/manboster/manboster/internal/engine/gatekeeper"
	"github.com/manboster/manboster/internal/engine/gateway"
	"github.com/manboster/manboster/internal/engine/onboard"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/session/selection"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/llm"
)

type Handler struct {
	repo                    repository.Repository
	llmProviders            map[string]llm.Provider
	onboard                 *onboard.Service
	chatDataService         *chatdata.Service
	toolMaps                map[string]tool.Provider
	gateway                 *gateway.Service
	selectionSessionManager *selection.Manager
	gatekeeperService       *gatekeeper.Service
}

func NewHandler(repo repository.Repository, llmProviders map[string]llm.Provider, chatDataService *chatdata.Service, onboardService *onboard.Service, toolMaps map[string]tool.Provider, gatewayService *gateway.Service, selectionSessionManager *selection.Manager, gatekeeperService *gatekeeper.Service) *Handler {
	return &Handler{
		repo:                    repo,
		llmProviders:            llmProviders,
		chatDataService:         chatDataService,
		onboard:                 onboardService,
		toolMaps:                toolMaps,
		gateway:                 gatewayService,
		selectionSessionManager: selectionSessionManager,
		gatekeeperService:       gatekeeperService,
	}
}
