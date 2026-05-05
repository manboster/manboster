package manboster

import (
	"fmt"
	"os"

	"github.com/manboster/manboster/internal/cli/manboster/app"
	"github.com/manboster/manboster/internal/cli/manboster/daemon"
	"github.com/manboster/manboster/internal/cli/manboster/interactive"
	"github.com/spf13/cobra"
)

// Init uses cobra to register as a client in Manboster.
func Init() {
	var rootCmd = &cobra.Command{
		Use:   "manboster",
		Short: "manboster: Your Personal Manbo Lobster",
		Long:  `Powered by LLM, manboster is an AI assistant delivers you the best experience while keeping your device fully secured!`,
		Run:   app.Main,
	}

	// Add version indicator.
	rootCmd.AddCommand(app.VersionCmd())

	// Add configuration options
	rootCmd.AddCommand(interactive.OnboardConfigCmd())
	rootCmd.AddCommand(interactive.ConfigCmd())

	// Add daemon options
	rootCmd.AddCommand(daemon.StartCmd())
	rootCmd.AddCommand(daemon.StopCmd())
	rootCmd.AddCommand(daemon.RestartCmd())
	rootCmd.AddCommand(daemon.StatusCmd())
	// Add log options (along with daemon)
	rootCmd.AddCommand(daemon.LogCmd())

	// Disable smart completion in order to clean help, no more about it!
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
