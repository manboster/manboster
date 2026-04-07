package config

import (
	"os"
	"path/filepath"
)

// Path gets the correct path of manboster file storage
func Path(filename string) string {
	var p string
	p, err := os.UserHomeDir()
	if err != nil {
		return filename
	}

	if _, err := os.ReadDir(filepath.Join(p, ".manboster")); err != nil {
		err = os.MkdirAll(filepath.Join(p, ".manboster"), 0700)
		if err != nil {
			return filename
		}
	}
	p = filepath.Join(p, ".manboster", filename)
	return p
}
