package loader

import (
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/database"
	"github.com/manboster/manboster/internal/engine"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/repository"
)

type Loader struct {
	db            *database.Client
	repo          repository.Repository
	cfg           *config.Config
	engine        *engine.Engine
	llmProviders  []llm.Provider
	chatProviders []chat.Provider
}

func New(cfg *config.Config) *Loader {
	return &Loader{
		cfg: cfg,
	}
}
