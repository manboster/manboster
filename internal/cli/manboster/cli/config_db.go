package cli

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/repository"
)

type databaseConfigService struct {
	repo repository.Repository
}

type databaseConfigLandingSelection string

const (
	databaseConfigLandingSelectionUser    databaseConfigLandingSelection = "user"
	databaseConfigLandingSelectionSession databaseConfigLandingSelection = "session"
	databaseConfigLandingSelectionSoul    databaseConfigLandingSelection = "soul"
	databaseConfigLandingSelectionQuit    databaseConfigLandingSelection = "quit"
)

func newDatabaseConfigService(repo repository.Repository) *databaseConfigService {
	return &databaseConfigService{repo}
}

func (s *databaseConfigService) configDatabaseLandingForm() error {
	for {
		se, err := s.runConfigDatabaseLandingSelection()
		if err != nil {
			return err
		}
		switch se {
		case databaseConfigLandingSelectionUser:
			return nil
		case databaseConfigLandingSelectionSession:
			return nil
		case databaseConfigLandingSelectionSoul:
			return nil
		case databaseConfigLandingSelectionQuit:
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
				huh.NewOption("Sessions\nIt's chat sessions you made in daily chatting routine, you can purge it to save disk or do more.", databaseConfigLandingSelectionSession),
				huh.NewOption("Users\nIt helps you to grant a user to an admin or degrade it. It's all up to you to decide.", databaseConfigLandingSelectionUser),
				huh.NewOption("Souls\nSouls is the key system prompt for you to personalize your Manbo Lobster.", databaseConfigLandingSelectionSoul),
				huh.NewOption("Quit Manboster Configuration Wizard", databaseConfigLandingSelectionQuit),
			).Title("Please select what to configure in database").Description("Please choose which field you want to configure in database field.").Value(&se),
		)).Run()
	if err != nil {
		return "", err
	}
	return se, nil
}
