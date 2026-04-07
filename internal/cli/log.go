package cli

import (
	"fmt"
	"io"

	"github.com/fatih/color"
	"github.com/hpcloud/tail"
	"github.com/manboster/manboster/internal/config"
	"github.com/spf13/cobra"
)

// logCmd initializes cobra command
func logCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "log",
		Short: "Get active logs from Manboster daemon",
		Run:   logCommandExecutor,
	}
}

// logCommandExecutor gets information about manboster daemon
func logCommandExecutor(cmd *cobra.Command, args []string) {
	logPath := config.Path("manboster.log")

	cfg := tail.Config{
		Follow:    true,                                            // following info
		ReOpen:    true,                                            // if rotates, reopens this file
		MustExist: false,                                           // file may exist
		Poll:      true,                                            // using poll to get log
		Location:  &tail.SeekInfo{Offset: 0, Whence: io.SeekStart}, // full info
	}

	t, err := tail.TailFile(logPath, cfg)
	if err != nil {
		color.Red(fmt.Sprintf("follow tail file unsuccessful, error: %q", err))
		return
	}

	color.Cyan(">>> Reading log file...")

	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
