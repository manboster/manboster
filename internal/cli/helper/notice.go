package helper

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"

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

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return
	}
}
