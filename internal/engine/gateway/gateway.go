package gateway

import (
	"github.com/manboster/manboster/internal/tool"
)

type Service struct {
	toolProviders []tool.Provider
}

func NewService(toolProviders []tool.Provider) *Service {
	return &Service{
		toolProviders: toolProviders,
	}
}
