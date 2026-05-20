package huh

import (
	"os"
	"os/exec"
	"runtime"

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

func getHeight() int {
	// Try stderr first (less likely to be redirected), then stdout, then stdin.
	for _, fd := range []uintptr{os.Stderr.Fd(), os.Stdout.Fd(), os.Stdin.Fd()} {
		_, h, err := term.GetSize(fd)
		if err == nil && h > 0 {
			// Reserve space for title, description, help line, borders and padding.
			height := h - 10
			if height < 3 {
				height = 3
			}
			return height
		}
	}
	return 10 // safe fallback
}
