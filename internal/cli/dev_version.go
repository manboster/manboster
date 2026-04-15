package cli

import "github.com/spf13/cobra"

// versionCmd Register Cobra Commands, giving user version info, simple.
func devVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Get the version of Manbodev",
		Run:   versionCmdExecutor,
	}
}
