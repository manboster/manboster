package engine

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/engine/chatdata"
	"github.com/manboster/manboster/internal/engine/onboard"
	"github.com/manboster/manboster/internal/engine/safeguard"
	"github.com/manboster/manboster/internal/engine/soul"
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

	e.chatDataService = chatdata.New(e.repo, e.sessionManager, e.llmProviders)
	e.safeguardService = safeguard.New(e.repo)
	e.soulService = soul.New(e.repo)

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
