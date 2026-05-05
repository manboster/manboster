package interactive

import (
	"context"
	"fmt"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/cli/helper"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/repository/types"
)

type databaseConfigService struct {
	repo     repository.Repository
	sessions []types.Session
	chats    []types.Chat
	chatsMap map[string][]types.Chat
}

type databaseConfigLandingSelection string

const (
	databaseConfigLandingUser    databaseConfigLandingSelection = "user"
	databaseConfigLandingSession databaseConfigLandingSelection = "session"
	databaseConfigLandingSoul    databaseConfigLandingSelection = "soul"
	databaseConfigLandingQuit    databaseConfigLandingSelection = "quit"
)

func newDatabaseConfigService(repo repository.Repository) *databaseConfigService {
	return &databaseConfigService{
		repo:     repo,
		chatsMap: make(map[string][]types.Chat),
	}
}

type databaseConfigSessionSelection string

const (
	databaseConfigSessionPurge  databaseConfigSessionSelection = "purge"
	databaseConfigSessionSelect databaseConfigSessionSelection = "select"
	databaseConfigSessionQuit   databaseConfigSessionSelection = "quit"
)

func (s *databaseConfigService) runConfigDatabaseSessionSelection(ctx context.Context) error {
	defer helper.ClearScreen()

	var se databaseConfigSessionSelection
	err := s.printConfigDatabaseSessionList(ctx)
	if err != nil {
		return err
	}
	err = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[databaseConfigSessionSelection]().Options(
				huh.NewOption("Purge Unused Session Data (Save Space)", databaseConfigSessionPurge),
				huh.NewOption("Select Sessions", databaseConfigSessionSelect),
				huh.NewOption("Quit", databaseConfigSessionQuit),
			).Value(&se).Description("What do you want to do?"),
		)).Run()
	if err != nil {
		return err
	}
	helper.ClearScreen()
	switch se {
	case databaseConfigSessionPurge:
		purgeNum := len(s.sessions) - len(s.chatsMap)
		if purgeNum <= 0 {
			color.Yellow("No need to purge sessions, your session list is clean and smart!")
			time.Sleep(500 * time.Millisecond)
			return nil
		}
		if helper.ContinueConfirm(ctx, fmt.Sprintf("Do you really want to DELETE %d unused sessions? YOUR ACTION IS IRREVERSIBLE!", purgeNum)) {
			err := s.purgeConfigDatabaseSession(ctx)
			if err != nil {
				return err
			}
			color.Blue(fmt.Sprintf("Successfully deleted %d unused sessions!", purgeNum))
			time.Sleep(500 * time.Millisecond)
		}
		return nil
	case databaseConfigSessionSelect:
		if len(s.sessions) == 0 {
			color.Yellow("No sessions found!")
			time.Sleep(500 * time.Millisecond)
			return nil
		}
		err := s.runConfigDatabaseSessionSelect(ctx)
		if err != nil {
			return err
		}
		return nil
	case databaseConfigSessionQuit:
		color.Blue("Bye!")
		return nil
	default:
		return fmt.Errorf("unexpected database session selection: %s", se)
	}
}
