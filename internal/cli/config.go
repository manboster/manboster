package cli

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/spf13/cobra"
)

func configCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config [command]",
		Short: "Run configuration wizard for Manboster application",
		Args:  cobra.MaximumNArgs(1),
		Run:   configCmdRun,
	}

	cmd.Flags().BoolP("edit", "e", false, "Open config file in terminal $EDITOR")

	editCmd := &cobra.Command{
		Use:   "edit",
		Short: "Open config file in terminal $EDITOR",
		Args:  cobra.NoArgs,
		Run:   configCmdEditRun,
	}
	openCmd := &cobra.Command{
		Use:   "open",
		Short: "Open config file using your system's default editor",
		Args:  cobra.NoArgs,
		Run:   configCmdOpenRun,
	}
	helpCmd := &cobra.Command{
		Use:   "help",
		Short: "The help of this subcommand",
		Args:  cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			if c.HasParent() {
				err := c.Parent().Help()
				if err != nil {
					return
				}
			}
		},
	}

	cmd.AddCommand(helpCmd)
	cmd.AddCommand(editCmd)
	cmd.AddCommand(openCmd)

	return cmd
}

// configCmdRun is used to run interactive huh forms to config.
func configCmdRun(cmd *cobra.Command, args []string) {
	err := configFormRun()
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Client] We encountered an error when configuring."))
	}
}

// configCmdEditRun runs terminal editor to config
func configCmdEditRun(cmd *cobra.Command, args []string) {
	err := config.Init()
	if err != nil {
		if errors.Is(err, config.ErrNoConfig) && runtime.GOOS != "windows" {
			color.Yellow(fmt.Sprintf("[Manboster Client] Config file not found!"))
			color.Yellow(fmt.Sprintf("Do you want to create a new one and edit it? [y/N]: "))

			var input string
			_, _ = fmt.Scanln(&input)
			if strings.ToLower(input) != "y" {
				color.Cyan("Operation cancelled. Please run `manboster onboard` for a guided setup.")
				return
			}
		} else {
			color.Red(fmt.Sprintf("[Manboster Client] Error initializing config: %q", err))
			return
		}
	}
	p := config.Path("config.yaml")
	err = openEditor(p)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Client] Error opening config file: %q", err))
		return
	}
}

// configCmdOpenRun
func configCmdOpenRun(cmd *cobra.Command, args []string) {
	err := config.Init()
	if err != nil {
		if errors.Is(err, config.ErrNoConfig) {
			color.Red(fmt.Sprintf("[Manboster Client] Config file not found! Please run `manboster onboard` at least once!"))
		} else {
			color.Red(fmt.Sprintf("[Manboster Client] Error initializing config: %q", err))
		}
		return
	}
	p := config.Path("config.yaml")
	err = openWithSystemDefault(p)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Client] Error opening config file: %q", err))
		return
	}
}

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

	// 这里不需要绑定 Stdin/Stdout，因为它是弹出一个独立的 GUI 窗口
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
