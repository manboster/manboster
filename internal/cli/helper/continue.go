package helper

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

func ContinueConfirm(ctx context.Context, content string) bool {
	agree := false

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(fmt.Sprintf("%s\nContinue?", content)).
				Affirmative("Continue").
				Negative("Skip").
				Value(&agree),
		),
	)

	err := form.Run()
	if err != nil {
		os.Exit(0)
	}
	return agree
}
