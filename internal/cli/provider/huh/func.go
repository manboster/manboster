package huh

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/x/term"
	"github.com/manboster/manboster/spec/cli"
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

func BuildHuhOptions(options []cli.Option) []huh.Option[string] {
	var ops []huh.Option[string]
	for _, o := range options {
		opt := huh.NewOption[string](o.Key, o.Value)
		if o.Selected {
			opt.Selected(o.Selected)
		}
		ops = append(ops, opt)
	}
	return ops
}

func BuildProviderOption(options []cli.Option, resp string) (cli.Option, bool) {
	for _, option := range options {
		if option.Value == resp {
			return option, true
		}
	}
	return cli.Option{}, false
}

func BuildProviderOptions(options []cli.Option, resp []string) []cli.Option {
	var ops []cli.Option
	for _, r := range resp {
		for _, o := range options {
			if o.Value == r {
				ops = append(ops, o)
				break
			}
		}
	}
	return ops
}

// sizeModel is a minimal bubbletea model that captures the terminal size
// from WindowSizeMsg and immediately quits.
type sizeModel struct{ h int }

func (m sizeModel) Init() tea.Cmd                            { return tea.WindowSize() }
func (m sizeModel) View() string                             { return "" }
func (m sizeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if sz, ok := msg.(tea.WindowSizeMsg); ok {
		m.h = sz.Height
		return m, tea.Quit
	}
	return m, nil
}

func getHeight() int {
	const reserve = 10

	// Try each stdio fd first — fast path for normal terminals.
	for _, fd := range []uintptr{os.Stderr.Fd(), os.Stdout.Fd(), os.Stdin.Fd()} {
		if _, h, err := term.GetSize(fd); err == nil && h > 0 {
			return max(h-reserve, 3)
		}
	}

	// Fallback: LINES env var (set by some terminals).
	if lines := os.Getenv("LINES"); lines != "" {
		if h, err := strconv.Atoi(lines); err == nil && h > 0 {
			return max(h-reserve, 3)
		}
	}

	// Last resort: run a minimal bubbletea program to receive WindowSizeMsg.
	// This works in JetBrains and other pty-emulated environments.
	if m, err := tea.NewProgram(sizeModel{}, tea.WithoutRenderer()).Run(); err == nil {
		if sm, ok := m.(sizeModel); ok && sm.h > 0 {
			return max(sm.h-reserve, 3)
		}
	}

	return 10
}
