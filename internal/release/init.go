package release

import (
	"runtime/debug"
	"strings"
)

func init() {
	if bi, avail := debug.ReadBuildInfo(); avail {
		if bi.Main.Version != "(devel)" && bi.Main.Version != "" {
			BuildCommit = "Go install, version " + bi.Main.Version
			Package = "goinstall"
		}

		// if invalid so we injected this
		for _, setting := range bi.Settings {
			switch setting.Key {
			case "vcs.revision":
				if BuildCommit == "unknown" || strings.HasPrefix(BuildCommit, "Go install") {
					BuildCommit = setting.Value[:6]
				}
			case "vcs.time":
				if BuildTime == "unknown" {
					BuildTime = setting.Value
				}
			}
		}
	}
}
