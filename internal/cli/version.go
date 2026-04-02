package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd Register Cobra Commands, giving user version info, simple.
func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Get the version of Manboster",
		Run:   versionCmdExecutor,
	}
}

func versionCmdExecutor(cmd *cobra.Command, args []string) {
	fmt.Println("0.0.0")
}
