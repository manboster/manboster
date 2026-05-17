package interact

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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
