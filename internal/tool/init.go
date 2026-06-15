package tool

import (
	"os"

	"github.com/manboster/manboster/internal/config"
)

var IsLoading = false

func init() {
	IsLoading = detectLoaderMode()
}

func detectLoaderMode() bool {
	if err := config.Init(); err != nil {
		return false
	}

	if len(os.Args) < 2 {
		return true
	}
	for _, arg := range os.Args[1:] {
		if len(arg) > 0 && arg[0] == '-' {
			continue
		}
		return arg != "onboard" && arg != "config"
	}

	return true
}
