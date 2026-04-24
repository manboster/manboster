package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Init uses cobra to register as a client in Manboster.
func Init() {
	var rootCmd = &cobra.Command{
		Use:   "manboster",
		Short: "manboster: Your Personal Manbo Lobster",
		Long:  `Powered by LLM, manboster is an AI assistant delivers you the best experience while keeping your device fully secured!`,
		Run:   main,
	}

	// Add version indicator.
	rootCmd.AddCommand(versionCmd())

	// Add configuration options
	rootCmd.AddCommand(onboardConfigCmd())
	rootCmd.AddCommand(configCmd())

	// Add daemon options
	rootCmd.AddCommand(startCmd())
	rootCmd.AddCommand(stopCmd())
	rootCmd.AddCommand(restartCmd())
	rootCmd.AddCommand(statusCmd())
	// Add log options (along with daemon)
	rootCmd.AddCommand(logCmd())

	// Disable smart completion in order to clean help, no more about it!
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
