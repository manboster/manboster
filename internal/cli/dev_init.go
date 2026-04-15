package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// DevInit uses cobra to register as a client in manbodev.
func DevInit() {
	var rootCmd = &cobra.Command{
		Use:   "manbodev",
		Short: "manbodev: A helpful plugin build toolchain for Manboster.",
		Long:  `Manboster is an build helper delivers you the best experience in developing your plugin for Manboster!`,
		Run:   devMain,
	}

	// Add version indicator.
	rootCmd.AddCommand(devVersionCmd())

	// Disable smart completion in order to clean help, no more about it!
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
