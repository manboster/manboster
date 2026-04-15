package cli

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/spf13/cobra"
)

// configCmd provides an interactive TUI to configure your manboster application.
func configCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Configurations for Manboster application",
		Run:   ConfigCmdRun,
	}
}

// ConfigCmdRun is used to run interactive huh forms.
func ConfigCmdRun(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	cfg, err := ConfigurationForm(ctx)
	if err != nil {
		os.Exit(1)
		return
	}
	err = config.Write(cfg)
	if err != nil {
		os.Exit(1)
		return
	}
	color.Green("Configuration writing successful!")
}
