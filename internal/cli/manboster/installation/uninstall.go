package installation

import "github.com/spf13/cobra"

func UninstallCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstall all Manboster components with data",
		Run:   uninstallCmd,
	}
}

func uninstallCmd(cmd *cobra.Command, args []string) {
	
}
