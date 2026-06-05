package manbofs

import (
	"os"

	"github.com/manboster/manboster/internal/util"
)

func Read(p string) (string, error) {
	realPath, err := util.SafePath(path, p)
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(realPath)
	if err != nil {
		return "", err
	}

	bytes, err := util.BytesToBase64URL(data)
	if err != nil {
		return "", err
	}

	return bytes, nil
}
