package manbodev

import (
	"fmt"
	"runtime"

	"github.com/manboster/manboster/internal/config"
	"github.com/spf13/cobra"
)

// versionCmd Register Cobra Commands, giving user version info, simple.
func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Get the version of Manbodev",
		Run:   versionCmdExecutor,
	}
}

func versionCmdExecutor(cmd *cobra.Command, args []string) {
	fmt.Printf("Manbodev version %s %s, commit %s, build at %s %s/%s\n", config.Version, config.CurrentVersion, config.BuildCommit, config.BuildTime, runtime.GOOS, runtime.GOARCH)
}
