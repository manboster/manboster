package config

import "github.com/charmbracelet/huh"

// Provider provides interfaces for all configurations
type Provider interface {
	Name() string
	DisplayName() string
	ToHuhGroup() []*huh.Group
	VerifyAndConvert() error
	GetConfig() any
}

// Config makes configurations to this whole application and store it in yml format.
type Config struct {
	Version int16        `yaml:"version"` // version data, used for migration.
	Chats   []ChatConfig `yaml:"chats"`   // chat configurations
	LLMs    []LLMConfig  `yaml:"llms"`    // LLM configurations
	App     AppConfig    `yaml:"app"`     // APP specific configurations
}

// ChatConfig stores configurations from chat providers
type ChatConfig struct {
	Provider      string `yaml:"provider"`      // Chat Provider Name
	Configuration any    `yaml:"configuration"` // Chat Provider specific configuration data
}

// LLMConfig stores configurations from LLM providers
type LLMConfig struct {
	Provider      string `yaml:"provider"` // LLM Provider Name
	Configuration any    `yaml:"configuration"`
}

// AppConfig stores configurations used for applications
type AppConfig struct {
}
