package helper

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
)

func ContinueConfirm(ctx context.Context, content string) bool {
	agree := false

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(fmt.Sprintf("%s\nContinue?", content)).
				Affirmative(i18n.T(keys.BtnContinue)).
				Negative(i18n.T(keys.BtnSkip)).
				Value(&agree),
		),
	)

	err := form.Run()
	if err != nil {
		os.Exit(0)
	}
	return agree
}
