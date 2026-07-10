package installation

import (
	"github.com/fatih/color"
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
}
