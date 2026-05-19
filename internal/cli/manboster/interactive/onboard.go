package interactive

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/spf13/cobra"
)

// OnboardConfigCmd provides an interactive TUI to configure your manboster application.
func OnboardConfigCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "onboard",
		Short: "Run onboard configuration wizard for Manboster application",
		Run:   OnboardConfigCmdRun,
	}
}

// OnboardConfigCmdRun is used to run interactive huh forms.
func OnboardConfigCmdRun(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	err := OnboardWarningForm(ctx)
	if err != nil {
		os.Exit(1)
		return
	}

	if config.VersionType(config.CurrentVersion) != config.VersionStable {
		err := OnboardVersionWarningForm(ctx)
		if err != nil {
			os.Exit(1)
			return
		}
	}

	cfg, err := OnboardConfigurationForm(ctx)
	if err != nil {
		os.Exit(1)
		return
	}

	err = config.Write(cfg, config.Path("config.yaml"))
	if err != nil {
		os.Exit(1)
		return
	}
	color.Green("Configuration writing successful!")
}
