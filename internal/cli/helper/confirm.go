package helper

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

func ConfirmForm(ctx context.Context, tips string, confirmTitle string, confirmContent string) error {
	agree := false

	width, _, _ := term.GetSize(os.Stdout.Fd())
	if width == 0 {
		width = 80
	}
	textWidth := width - 8

	renderer, err := glamour.NewTermRenderer(
		glamour.WithEnvironmentConfig(),
		glamour.WithWordWrap(textWidth),
	)
	if err != nil {
		panic(err)
	}
	renderedMD, _ := renderer.Render(tips)
	defer func(renderer *glamour.TermRenderer) {
		err := renderer.Close()
		if err != nil {
			panic(err)
		}
	}(renderer)

	boxStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(0, 2).
		MarginTop(1).
		MarginBottom(1)

	fmt.Println(boxStyle.Render(renderedMD))
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(confirmTitle).
				Affirmative(confirmContent).
				Negative("Exit Now").
				Value(&agree),
		),
	)
	err = form.Run()
	if !agree {
		os.Exit(0)
	}
	ClearScreen()
	return err
}
