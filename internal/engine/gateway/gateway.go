package gateway

import (
	"github.com/manboster/manboster/internal/session/selection"
	"github.com/manboster/manboster/internal/tool"
)

type Service struct {
	toolProviders           []tool.Provider
	selectionSessionManager *selection.Manager
}

func NewService(toolProviders []tool.Provider, selectionSessionManager *selection.Manager) *Service {
	return &Service{
		toolProviders:           toolProviders,
		selectionSessionManager: selectionSessionManager,
	}
}
