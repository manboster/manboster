package helper

import (
	"context"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
)

func ConfirmForm(ctx context.Context, tips string, confirmTitle string, confirmContent string) error {
	agree := false

	DisplayText(tips)
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(confirmTitle).
				Affirmative(confirmContent).
				Negative(i18n.T(keys.BtnExitNow)).
				Value(&agree),
		),
	)
	err := form.Run()
	if !agree {
		os.Exit(0)
	}
	ClearScreen()
	return err
}
