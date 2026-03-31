package cli

import (
	"fmt"
	"os"

	"github.com/manboster/manboster/core/config"
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
	cfg, err := config.Form()
	if err != nil {
		os.Exit(1)
		return
	}
	err = config.Write(cfg)
	fmt.Println("Successfully created config.yaml, open Manboster directly and enjoy it!")
	if err != nil {
		os.Exit(1)
		return
	}
}
