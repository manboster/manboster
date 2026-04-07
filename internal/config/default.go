package config

// Default sets Default configurations
func Default(c Config) Config {
	// write database default path
	c.App.DBPath = Path("manboster.db")

	// write current configuration version
	c.Version = V

	return c
}
