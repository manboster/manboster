package config

import "github.com/manboster/manboster/internal/release"

// Default sets Default configurations
func (c Config) Default() Config {
	// write database default path
	c.App.DBPath = Path("manboster.db")

	// write current configuration version
	c.Version = release.V

	return c
}
