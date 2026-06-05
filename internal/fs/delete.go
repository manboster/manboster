package manbofs

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/util"
)

func Delete(p string) error {
	realPath, err := util.SafePath(path, p)
	if err != nil {
		return err
	}
	err = os.Remove(realPath)
	if err != nil {
		return err
	}
	return nil
}

func Deletes(p []string) error {
	for _, p := range p {
		err := Delete(p)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Telegram Provider] Failed to delete: %v", err))
		}
	}
	return nil
}
