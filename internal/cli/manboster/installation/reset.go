package installation

import (
	"errors"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/cli/helper"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/spf13/cobra"
)

func ResetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "reset",
		Short: "Reset all Manboster components",
		Run:   resetCmd,
	}
}

func resetCmd(cmd *cobra.Command, args []string) {
	color.Red(i18n.T(keys.AppResetPrompt))
	chr := helper.GetChar()
	if strings.ToLower(chr) == "y" {
		err := reset()
		if err != nil {
			color.Green(i18n.Te(keys.AppResetError, "", err))
			os.Exit(1)
		}
		color.Red(i18n.T(keys.AppResetSuccess))
		os.Exit(0)
	}

	color.Yellow(i18n.Te(keys.AppResetError, "", errors.New("user canceled")))
	os.Exit(1)
}

func reset() error {
	location := config.Path("")

	return os.RemoveAll(location)
}
