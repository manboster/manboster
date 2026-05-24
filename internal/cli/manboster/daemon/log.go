package daemon

import (
	"fmt"
	"io"

	"github.com/fatih/color"
	"github.com/hpcloud/tail"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/spf13/cobra"
)

// LogCmd initializes cobra command
func LogCmd() *cobra.Command {
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
		Follow:    true,
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Location:  &tail.SeekInfo{Offset: 200, Whence: io.SeekStart},
	}

	t, err := tail.TailFile(logPath, cfg)
	if err != nil {
		color.Red(fmt.Sprintf(i18n.T(keys.DaemonLogError), err))
		return
	}

	color.Cyan(i18n.T(keys.DaemonLogReading))

	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
