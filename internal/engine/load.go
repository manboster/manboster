package engine

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/engine/chatdata"
	"github.com/manboster/manboster/internal/engine/command"
	"github.com/manboster/manboster/internal/engine/gatekeeper"
	"github.com/manboster/manboster/internal/engine/gateway"
	"github.com/manboster/manboster/internal/engine/handler"
	"github.com/manboster/manboster/internal/engine/onboard"
	"github.com/manboster/manboster/internal/engine/processor"
	"github.com/manboster/manboster/internal/engine/runner"
	"github.com/manboster/manboster/internal/engine/safeguard"
	"github.com/manboster/manboster/internal/engine/soul"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/session"
)

func (e *Engine) Load(ctx context.Context) error {
	// initialize variables
	color.Blue("[Manboster Engine] Loading engine...")

	// First, we get user counts using cache
	count, err := e.repo.UserCounts(ctx)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while getting user counts from repository, error: %s", err))
	}
	if err != nil || count == 0 {
		e.onboard = onboard.New()
	}

	e.soulService = soul.NewService(e.repo)
	e.sessionService = session.NewService(e.repo, e.soulService, e.config)

	e.safeguardService = safeguard.NewService(e.repo)

	err = e.soulService.Init(ctx)
	if err != nil {
		color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to initialize soul service: %q", err))
	}

	for _, tool := range e.toolProviders {
		_, avail := e.toolMaps[tool.Name()]
		if !avail {
			e.toolMaps[tool.Name()] = tool
		} else {
			fmt.Printf("[Manboster Engine] Duplicate tool '%s'! The next one will be ignored.\n", tool.Name())
		}
	}
	e.gateway = gateway.NewService(e.toolProviders, e.sessionService.Manager.SelectionManager)
	e.chatDataService = chatdata.NewService(e.repo, e.sessionService.Manager.ChatSession, e.sessionService, e.llmProviders, e.gateway)
	e.gatekeeperService = gatekeeper.NewService(e.gateway, e.safeguardService, e.config.Hachimi, e.hachimiProvider, e.hachimiLoaded, e.sessionService, e.llmProviders)
	e.handler = handler.NewHandler(e.repo, e.llmProviders, e.chatDataService, e.onboard, e.toolMaps, e.gateway, e.sessionService.Manager, e.gatekeeperService, e.safeguardService)
	e.commandHandler = command.NewHandler(e, e.repo, e.safeguardService, e.sessionService, e.llmProviders, e.config, e.soulService, e.onboard, e.handler)

	e.processor = processor.New(e, e.sessionService, e.safeguardService, e.onboard)
	runner.Instance = runner.NewRunner(e.processor, e.chatProviders)
	go func() {
		err := runner.Instance.Run(ctx)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] Failed to run engine runner: %q", err))
		}
	}()

	// version tips
	if config.ChannelType(config.CurrentChannel) != config.ChannelStable {
		color.Yellow(i18n.T(keys.EngineLoadUnstable))
		switch config.ChannelType(config.CurrentChannel) {
		case config.ChannelRC:
			color.Yellow(i18n.T(keys.EngineLoadRC))
		case config.ChannelBeta:
			color.Yellow(i18n.T(keys.EngineLoadBeta))
		case config.ChannelAlpha:
			color.Yellow(i18n.T(keys.EngineLoadAlpha))
		case config.ChannelCanary, config.ChannelNightly:
			color.HiRed(i18n.T(keys.EngineLoadCanary))
			color.Yellow(i18n.T(keys.EngineLoadCanaryWarn))
		default:
		}
	}

	return nil
}
