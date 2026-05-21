package config

import (
	"github.com/fatih/color"
)

func (c Config) Validate() error {
	current := int16(0)
	// check version
	if c.Version > current {
		color.Yellow("Configuration contains an unsupported version, if you want to use this configuration, please download the latest version. Or you can reconfigure it with `manboster onboard`.")
		return ErrInvalidConfig
	}
	if c.Version < current {
		color.Yellow("Outdated configuration, if you want to use this configuration, please run `manboster upgrade` to upgrade your old data. Or you can reconfigure it with `manboster onboard`.")
		return ErrInvalidConfig
	}

	// check valid configuration
	if len(c.Chats) == 0 {
		color.Red("Missing chat configuration, please edit it with `manboster config` or reconfigure it with `manboster onboard`.")
		return ErrInvalidConfig
	}
	if len(c.LLMs) == 0 {
		color.Red("Missing LLM configuration, please edit it with `manboster config` or reconfigure it with `manboster onboard`.")
		return ErrInvalidConfig
	}
	if c.Hachimi.Enabled && len(c.Hachimi.Hachimi) == 0 {
		color.Red("Hachimi enabled but there is available providers!")
		return ErrInvalidConfig
	}
	if c.App.DBPath == "" {
		color.Red("Missing Database Path! If it's not valid, please consider to use `manboster onboard` to reconfigure this app.")
		return ErrInvalidConfig
	}
	if c.App.DefaultLLMProvider == "" {
		color.Red("Missing Default LLM Provider! If it's not valid, please consider to use `manboster onboard` to reconfigure this app.")
		return ErrInvalidConfig
	}
	if c.App.DefaultLLMModel == "" {
		color.Red("Missing Default LLM Model! If it's not valid, please consider to use `manboster onboard` to reconfigure this app.")
		return ErrInvalidConfig
	}

	return nil
}
