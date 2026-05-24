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
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/spf13/cobra"
)

// configCmdRun is used to run interactive huh forms to config.
func configCmdRun(cmd *cobra.Command, args []string) {
	res, err := ctx.DaemonCtx.Search()
	if err == nil && res != nil {
		color.Red(i18n.T(keys.ConfigRunDaemonRunning))
		return
	}

	err = config.Init()
	if err != nil {
		color.Red(fmt.Sprintf(i18n.T(keys.ConfigRunInitError), err))
		return
	}

	err = runConfigEntrypoint(huh.Huh{})
	if err != nil {
		color.Red(fmt.Sprintf(i18n.T(keys.ConfigRunError), err))
	}
}

// configCmdEditRun runs terminal editor to config
func configCmdEditRun(cmd *cobra.Command, args []string) {
	err := config.Init()
	if err != nil {
		if errors.Is(err, config.ErrNoConfig) && runtime.GOOS != "windows" {
			color.Yellow(i18n.T(keys.ConfigEditNotFound))
			color.Yellow(i18n.T(keys.ConfigEditCreatePrompt))

			var input string
			_, _ = fmt.Scanln(&input)
			if strings.ToLower(input) != "y" {
				color.Cyan(i18n.T(keys.ConfigEditCancelled))
				return
			}
		} else {
			color.Red(fmt.Sprintf(i18n.T(keys.ConfigRunInitError), err))
			return
		}
	}
	p := config.Path("config.yaml")
	err = openEditor(p)
	if err != nil {
		color.Red(fmt.Sprintf(i18n.T(keys.ConfigEditOpenError), err))
		return
	}
}

// configCmdOpenRun
func configCmdOpenRun(cmd *cobra.Command, args []string) {
	err := config.Init()
	if err != nil {
		if errors.Is(err, config.ErrNoConfig) {
			color.Red(i18n.T(keys.ConfigOpenNotFound))
		} else {
			color.Red(fmt.Sprintf(i18n.T(keys.ConfigRunInitError), err))
		}
		return
	}
	p := config.Path("config.yaml")
	err = openWithSystemDefault(p)
	if err != nil {
		color.Red(fmt.Sprintf(i18n.T(keys.ConfigOpenError), err))
		return
	}
}

// OnboardConfigCmdRun is used to run interactive huh forms.
func OnboardConfigCmdRun(cmd *cobra.Command, args []string) {
	cfg, err := runOnboardConfig(huh.Huh{})
	if err != nil {
		color.Red(fmt.Sprintf(i18n.T(keys.ConfigOnboardError), err))
		os.Exit(1)
		return
	}

	err = cfg.Validate()
	if err != nil {
		color.Red(fmt.Sprintf(i18n.T(keys.ConfigOnboardValidateErr), err))
		return
	}
	err = config.Write(cfg, config.Path("config.yaml"))
	if err != nil {
		color.Red(fmt.Sprintf(i18n.T(keys.ConfigOnboardWriteError), err))
		os.Exit(1)
		return
	}
	color.Green(i18n.T(keys.ConfigOnboardSuccess))
}
