package engine

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/engine/chatdata"
	"github.com/manboster/manboster/internal/engine/command"
	"github.com/manboster/manboster/internal/engine/gatekeeper"
	"github.com/manboster/manboster/internal/engine/gateway"
	"github.com/manboster/manboster/internal/engine/handler"
	"github.com/manboster/manboster/internal/engine/onboard"
	"github.com/manboster/manboster/internal/engine/processor"
	"github.com/manboster/manboster/internal/engine/safeguard"
	"github.com/manboster/manboster/internal/engine/soul"
	"github.com/manboster/manboster/internal/hachimi"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/session"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

type Engine struct {
	sessionService  *session.Service
	llmProviders    map[string]llm.Provider
	toolProviders   []tool.Provider
	chatProviders   map[string]chat.Provider
	hachimiProvider hachimi.Provider
	hachimiLoaded   *bool
	toolMaps        map[string]tool.Provider
	config          *config.Config
	repo            repository.Repository

	commandHandler    *command.Handler
	handler           *handler.Handler
	gateway           *gateway.Service
	onboard           *onboard.Service
	safeguardService  *safeguard.Service
	chatDataService   *chatdata.Service
	soulService       *soul.Service
	gatekeeperService *gatekeeper.Service
	processor         *processor.Service
}

func New(cfg *config.Config, repo repository.Repository, llmProviders map[string]llm.Provider, chatProviders map[string]chat.Provider, toolProviders []tool.Provider, hachimiProvider hachimi.Provider, hachimiLoaded *bool) (*Engine, error) {
	return &Engine{
		llmProviders:    llmProviders,
		toolProviders:   toolProviders,
		toolMaps:        make(map[string]tool.Provider),
		chatProviders:   chatProviders,
		config:          cfg,
		repo:            repo,
		onboard:         nil,
		hachimiLoaded:   hachimiLoaded,
		hachimiProvider: hachimiProvider,
	}, nil
}
