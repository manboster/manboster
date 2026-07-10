package _30

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/migration/patches"
)

var cacheMigrationPatch = patches.Patch{
	Detect: func(conf config.Config) bool {
		return config.DetectOlderPath()
	},
	Migrate: func(conf config.Config) (config.Config, error) {
		_ = godotenv.Load()
		dir := os.Getenv("MANBOSTER_HOME")
		p := ""
		if dir == "" {
			p, _ = os.UserHomeDir()
		} else {
			p = dir
		}

		oldPath := filepath.Join(p, ".manboster")
		newPath := config.Path("")

		err := os.Rename(oldPath, newPath)
		if err != nil {
			return conf, err
		}
		return conf, nil
	},
	Description: i18n.T(keys.Migration030HomeDirDetected),
	Introduce:   "0.3.0",
	V:           0,
}
