package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Path gets the correct path of manboster file storage
func Path(filename string) string {
	_ = godotenv.Load()
	dir := os.Getenv("MANBOSTER_HOME")

	var p string
	if dir == "" {
		var err error
		p, err = os.UserHomeDir()
		if err != nil {
			return filename
		}
	} else {
		p = dir
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
