package interact

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/cli/manboster/ctx"
	"github.com/manboster/manboster/internal/cli/provider/huh"
	"github.com/manboster/manboster/internal/config"
	"github.com/spf13/cobra"
)

// configCmdRun is used to run interactive huh forms to config.
func configCmdRun(cmd *cobra.Command, args []string) {
	res, err := ctx.DaemonCtx.Search()
	if err == nil && res != nil {
		color.Red("Manboster is running, please run 'manboster stop' to stop it!\nQuiting the application...")
		return
	}
	// err = configFormRun()
	// if err != nil {
	// 	color.Red(fmt.Sprintf("[Manboster Client] We encountered an error when configuring: %q.", err))
	// }
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

// OnboardConfigCmdRun is used to run interactive huh forms.
func OnboardConfigCmdRun(cmd *cobra.Command, args []string) {
	cfg, err := runOnboardConfig(huh.Huh{})
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Client] Error while configuring: %q", err))
		os.Exit(1)
		return
	}

	err = cfg.Validate()
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Client] Error while validating configuration: %q", err))
		return
	}
	err = config.Write(cfg, config.Path("config.yaml"))
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Client] Error while writing config: %q", err))
		os.Exit(1)
		return
	}
	color.Green("Configuration writing successful!")
}
