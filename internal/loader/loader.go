package loader

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/database"
	"github.com/manboster/manboster/internal/engine"
	"github.com/manboster/manboster/internal/hachimi"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

type Loader struct {
	db              *database.Client
	repo            repository.Repository
	cfg             *config.Config
	engine          *engine.Engine
	llmProviders    map[string]llm.Provider
	toolProviders   map[string]tool.Provider
	chatProviders   []chat.Provider
	hachimiProvider hachimi.Provider
}

func New(cfg *config.Config) *Loader {
	return &Loader{
		cfg: cfg,
	}
}
