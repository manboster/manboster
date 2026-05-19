package huh

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

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

func (h Huh) Display(content string, timeout time.Duration) error {
	ClearScreen()
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
	renderedMD, _ := renderer.Render(content)
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

	if timeout > 0 {
		defer func() {
			time.Sleep(timeout)
			ClearScreen()
		}()
	}

	return nil
}
