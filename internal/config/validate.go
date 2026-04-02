package config

import (
	"github.com/fatih/color"
)

func (c *Config) Validate() error {
	current := int16(0)
	// check version
	if c.Version > current {
		color.Yellow("Configuration contains an unsupported version, if you want to use this configuration, please download the latest version. Or you can reconfigure it with `manboster config`.")
		return ErrInvalidConfig
	}
	if c.Version < current {
		color.Yellow("Outdated configuration, if you want to use this configuration, please run `manboster upgrade` to upgrade your old data. Or you can reconfigure it with `manboster config`.")
		return ErrInvalidConfig
	}

	// check valid configuration
	if len(c.Chats) == 0 {
		color.Red("Missing chat configuration, please reconfigure it with `manboster config`.")
		return ErrInvalidConfig
	}
	if len(c.LLMs) == 0 {
		color.Red("Missing LLM configuration, please reconfigure it with `manboster config`.")
		return ErrInvalidConfig
	}

	return nil
}
