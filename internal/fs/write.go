package manbofs

import (
	"io"
	"os"

	"github.com/manboster/manboster/internal/util"
)

func Write(p string, reader io.Reader) error {
	realPath, err := util.SafePath(path, p)
	if err != nil {
		return err
	}

	bytes, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	err = os.WriteFile(realPath, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
