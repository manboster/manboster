package interact

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/manboster/manboster/spec/cli"
	"github.com/manboster/manboster/spec/config"
)

// openWithSystemDefault opens a file with default open way
func openWithSystemDefault(filePath string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin": // macOS
		cmd = exec.Command("open", filePath)
	case "windows": // windows
		cmd = exec.Command("cmd", "/c", "start", filePath)
	case "linux": // Linux
		cmd = exec.Command("xdg-open", filePath)
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Run()
}

// openEditor opens default terminal editor for the user.
func openEditor(filePath string) error {
	// check out the default terminal editor
	editor := "vim"
	e := os.Getenv("EDITOR")
	if e != "" {
		editor = e
	}
	if runtime.GOOS == "windows" {
		editor = "notepad"
	}

	cmd := exec.Command(editor, filePath)

	// bind IO stream
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// run
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to open editor %s: %w", editor, err)
	}

	return nil
}

func buildOptions(p []config.Provider, selected []string) []cli.Option {
	var options []cli.Option
	for i, provider := range p {
		option := cli.Option{
			Key:   provider.DisplayName(),
			Value: provider.Name(),
		}

		if i == 0 && (len(selected) == 0 || selected == nil) {
			option.Selected = true
		} else if selected != nil && len(selected) > 0 {
			for j, selectedItem := range selected {
				if selectedItem == provider.Name() {
					option.Selected = true
					selected = append(selected[:j], selected[j+1:]...)
					break
				}
			}
		}

		options = append(options, option)
	}
	return options
}
