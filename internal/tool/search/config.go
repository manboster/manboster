package search

import "github.com/manboster/manboster/spec/config"

type Config struct {
}

func (c *Config) Name() string {
	return metadata.Name
}

func (c *Config) DisplayName() string {
	return metadata.DisplayName
}

func (c *Config) Args() *config.Args {
	return config.ArgsFromStruct(Config{})
}

func (c *Config) Validate() error {
	return nil
}

func (c *Config) GetConfig() any {
	return c
}
