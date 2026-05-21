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

	e.chatDataService = chatdata.NewService(e.repo, e.sessionService.Manager.ChatSession, e.llmProviders)
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
	e.gatekeeperService = gatekeeper.NewService(e.gateway, e.safeguardService, e.sessionService.Manager.Ignorance, e.config.Hachimi, e.hachimiProvider, e.hachimiLoaded)
	e.handler = handler.NewHandler(e.repo, e.llmProviders, e.chatDataService, e.onboard, e.toolMaps, e.gateway, e.sessionService.Manager, e.gatekeeperService, e.safeguardService)
	e.commandHandler = command.NewHandler(e.repo, e.safeguardService, e.sessionService, e.llmProviders, e.config, e.soulService, e.onboard, e.handler)

	e.processor = processor.New(e)

	runner.Instance = runner.NewRunner(e, e.chatProviders)
	go func() {
		err := runner.Instance.Run(ctx)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] Failed to run engine runner: %q", err))
		}
	}()

	// version tips
	if config.VersionType(config.CurrentVersion) != config.VersionStable {
		color.Yellow("[Manboster Engine] It seemed that you're using unstable version.")
		switch config.VersionType(config.CurrentVersion) {
		case config.VersionRC:
			color.Yellow("[Manboster Engine] You're using a Release Candidate build of Manboster! This version still has minor bugs to fix and we will appreciate if you report bugs you encountered. Sit tight and wait for the stable version!")
		case config.VersionBeta:
			color.Yellow("[Manboster Engine] You're using a Beta build of Manboster! This version marks new features have been landed and there is still something to be done. We will appreciate if you report bugs you encountered to us.")
		case config.VersionAlpha:
			color.Yellow("[Manboster Engine] You're using an Alpha build of Manboster! This version means new features is experimenting and there is a long way to implement. We will appreciate if you report bugs you encountered to us.")
		case config.VersionCanary:
			color.HiRed("[Manboster Engine] You're using a Canary build of Manboster! We apperiate your spirit to try something new, but please do not place any important information in this application!")
			color.Yellow("[Manboster Engine] Since the version is not ready to release, it's normal to have bugs. If you find any bug in this release, you can report to us in issues. Also, don't forget to check whether this bug is fixed or not!")
		default:
		}
	}

	return nil
}
