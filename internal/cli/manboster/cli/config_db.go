package cli

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

func (s *databaseConfigService) configDatabaseLandingForm() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		se, err := s.runConfigDatabaseLandingSelection()
		if err != nil {
			return err
		}
		switch se {
		case databaseConfigLandingUser:
			return nil
		case databaseConfigLandingSession:
			err := s.runConfigDatabaseSessionSelection(ctx)
			if err != nil {
				return err
			}
			return nil
		case databaseConfigLandingSoul:
			return nil
		case databaseConfigLandingQuit:
			color.Blue("Bye!")
			return nil
		default:
			return fmt.Errorf("unexpected database landing form: %s", se)
		}
	}
}

func (s *databaseConfigService) runConfigDatabaseLandingSelection() (databaseConfigLandingSelection, error) {
	var se databaseConfigLandingSelection
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[databaseConfigLandingSelection]().Options(
				huh.NewOption("Sessions\nIt's chat sessions you made in daily chatting routine, you can purge it to save disk or do more.", databaseConfigLandingSession),
				huh.NewOption("Users\nIt helps you to grant a user to an admin or degrade it. It's all up to you to decide.", databaseConfigLandingUser),
				huh.NewOption("Souls\nSouls is the key system prompt for you to personalize your Manbo Lobster.", databaseConfigLandingSoul),
				huh.NewOption("Quit Manboster Configuration Wizard", databaseConfigLandingQuit),
			).Title("Please select what to configure in database").Description("Please choose which field you want to configure in database field.").Value(&se),
		)).Run()
	if err != nil {
		return "", err
	}
	return se, nil
}

type databaseConfigSessionSelection string

const (
	databaseConfigSessionPurge  databaseConfigSessionSelection = "purge"
	databaseConfigSessionSelect databaseConfigSessionSelection = "select"
	databaseConfigSessionQuit   databaseConfigSessionSelection = "quit"
)

func (s *databaseConfigService) runConfigDatabaseSessionSelection(ctx context.Context) error {
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
			helper.ClearScreen()
			return nil
		}
		if helper.ContinueConfirm(ctx, fmt.Sprintf("Do you really want to DELETE %d unused sessions? YOUR ACTION IS IRREVERSIBLE!", purgeNum)) {
			err := s.purgeConfigDatabaseSession(ctx)
			if err != nil {
				return err
			}
			color.Blue(fmt.Sprintf("Successfully deleted %d unused sessions!", purgeNum))
			time.Sleep(500 * time.Millisecond)
			helper.ClearScreen()
		}
		return nil
	case databaseConfigSessionSelect:
		return nil
	case databaseConfigSessionQuit:
		color.Blue("Bye!")
		return nil
	default:
		return fmt.Errorf("unexpected database session selection: %s", se)
	}
}
