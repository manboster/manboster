package app

import (
	"fmt"
	"runtime"

	"github.com/manboster/manboster/internal/release"
	"github.com/spf13/cobra"
)

// VersionCmd Register Cobra Commands, giving user version info, simple.
func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Get the version of Manboster",
		Run:   versionCmdExecutor,
	}
}

func versionCmdExecutor(cmd *cobra.Command, args []string) {
	fmt.Printf("Manboster version %s %s, commit %s, build at %s %s/%s\n", release.Version, release.CurrentChannel, release.BuildCommit, release.BuildTime, runtime.GOOS, runtime.GOARCH)
}
