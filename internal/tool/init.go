package tool

import "os"

var IsLoading = false

func init() {
	IsLoading = detectLoaderMode()
}

func detectLoaderMode() bool {
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
