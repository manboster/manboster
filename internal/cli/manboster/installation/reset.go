package installation

import "github.com/spf13/cobra"

func ResetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "reset",
		Short: "Reset all Manboster components",
		Run:   resetCmd,
	}
}

func resetCmd(cmd *cobra.Command, args []string) {

}
