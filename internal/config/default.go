package config

import (
	"os"
	"path/filepath"
)

// Default sets Default configurations
func Default(c Config) Config {
	// write database default path
	dir, err := os.UserHomeDir()
	if err != nil {
		c.App.DBPath = "manboster.db"
	} else {
		c.App.DBPath = filepath.Join(dir, ".manboster", "manboster.db")
	}

	// write current configuration version
	c.Version = V

	return c
}
