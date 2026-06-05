package manbofs

import (
	"io"
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
