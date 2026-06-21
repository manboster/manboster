package config

import (
	"runtime/debug"
	"strings"
)

// Version defines manboster's application version.
const Version = "0.2.0"

// APILevel defines the current level(Tool) supported in Manboster.
const APILevel = 1

// V indicates config's version, now is 0
const V = 0

var (
	BuildCommit    string = "unknown"
	BuildTime      string = "unknown"
	CurrentChannel        = "unknown"
)

func init() {
	if bi, avail := debug.ReadBuildInfo(); avail {
		if bi.Main.Version != "(devel)" && bi.Main.Version != "" {
			BuildCommit = "Go install, version " + bi.Main.Version
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
