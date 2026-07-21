package _30

import (
	"fmt"
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

		// load legacy configurations
		err := config.Load(oldPath)
		if err != nil {
			return conf, err
		}
		conf = config.Read()

		oldInfo, err := os.Stat(oldPath)
		if os.IsNotExist(err) {
			return conf, nil
		}
		if err != nil {
			return conf, err
		}

		newInfo, err := os.Stat(newPath)
		if err == nil {
			if oldInfo.IsDir() && newInfo.IsDir() {
				entries, err := os.ReadDir(oldPath)
				if err != nil {
					return conf, err
				}
				for _, entry := range entries {
					oldItem := filepath.Join(oldPath, entry.Name())
					newItem := filepath.Join(newPath, entry.Name())

					if _, err := os.Stat(newItem); err == nil {
						backupItem := newItem + ".migration.bak"
						if err := os.Rename(newItem, backupItem); err != nil {
							return conf, fmt.Errorf("backup failed: %w", err)
						}
					}
					if err := os.Rename(oldItem, newItem); err != nil {
						return conf, fmt.Errorf("move %s failed: %w", entry.Name(), err)
					}
				}

				err = os.Remove(oldPath)
				if err != nil {
					return conf, err
				}
				conf.App.DBPath = filepath.Join(newPath, "manboster.db")
				return conf, nil
			}

			return conf, fmt.Errorf("cannot migrate: destination %s already exists with different type", newPath)
		}

		return conf, nil
	},
	Description: i18n.T(keys.Migration030HomeDirDetected),
	Introduce:   "0.3.0",
	V:           0,
}
