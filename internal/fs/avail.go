package manbofs

import (
	"os"

	"github.com/manboster/manboster/internal/util"
)

func Avail(p string) bool {
	realPath, err := util.SafePath(path, p)
	if err != nil {
		return false
	}
	if _, err := os.Stat(realPath); err != nil {
		return false
	}

	return true
}
