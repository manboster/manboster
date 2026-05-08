package interactive

import (
	"context"
	"fmt"

	"github.com/fatih/color"
)

func configLandingHachimiActionForm() (configLandingActionSelection, error) {
	return configLandingActionForm("Add a new hachimi provider", "Select an existing hachimi provider", "Quit")
}

func runLandingHachimiActionForm(ctx context.Context) error {
	printConfigHachimiProvidersData(ctx)
	se, err := configLandingChatActionForm()
	if err != nil {
		return err
	}
	switch se {
	case configLandingActionAdd:
	case configLandingActionSelect:
	case configLandingActionQuit:
		color.Blue("Bye!")
		return nil
	default:
		return fmt.Errorf("unexpected landing hachimi-action form: %s", se)
	}
	return nil
}
