package manboster

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/gofrs/flock"
	"github.com/manboster/manboster/internal/cli/manboster/app"
	"github.com/manboster/manboster/internal/cli/manboster/daemon"
	"github.com/manboster/manboster/internal/cli/manboster/interact"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
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
	rootCmd.AddCommand(interact.OnboardConfigCmd())
	rootCmd.AddCommand(interact.ConfigCmd())

	// Add daemon options
	rootCmd.AddCommand(daemon.StartCmd())
	rootCmd.AddCommand(daemon.StopCmd())
	rootCmd.AddCommand(daemon.RestartCmd())
	rootCmd.AddCommand(daemon.StatusCmd())
	// Add log options (along with daemon)
	rootCmd.AddCommand(daemon.LogCmd())

	// Disable smart completion in order to clean help, no more about it!
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// main inner check lock file is avail or not?
	lockPath := config.Path("manboster.lock")
	fileLock := flock.New(lockPath)
	defer func(fileLock *flock.Flock) {
		err := fileLock.Unlock()
		if err != nil {
			color.Yellow("[Manboster Client] Unlock error: %v", err)
		}
	}(fileLock)

	locked, err := fileLock.TryLock()
	if err != nil {
		color.Red("[Manboster Client] Lock error: %v", err)
		os.Exit(1)
	}

	if !locked {
		color.Red(i18n.T(keys.AppAnotherRunning))
		color.Red(i18n.T(keys.AppDaemonRunningQuit))
		os.Exit(1)
	}

	err = rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
