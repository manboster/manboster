package config

// Config makes configurations to this whole application and store it in yml format.
type Config struct {
	Version int16        `yaml:"version" mapstructure:"version" json:"version"` // version data, used for migration.
	Chats   []ChatConfig `yaml:"chats" mapstructure:"chats" json:"chats"`       // chat configurations
	LLMs    []LLMConfig  `yaml:"llms" mapstructure:"llms" json:"llms"`          // LLM configurations
	App     AppConfig    `yaml:"app" mapstructure:"app" json:"app"`             // APP specific configurations
}

// ChatConfig stores configurations from chat providers
type ChatConfig struct {
	Provider      string `yaml:"provider" mapstructure:"provider" json:"provider"`                // Chat Provider Name
	Configuration any    `yaml:"configuration" mapstructure:"configuration" json:"configuration"` // Chat Provider specific configuration data
}

// LLMConfig stores configurations from LLM providers
type LLMConfig struct {
	Provider      string `yaml:"provider" mapstructure:"provider" json:"provider"` // LLM Provider Name
	Configuration any    `yaml:"configuration" mapstructure:"configuration" json:"configuration"`
}

// AppConfig stores configurations used for applications
type AppConfig struct {
	DBPath string `yaml:"dbpath" mapstructure:"dbpath" json:"dbpath"` // SQLite's path
}

// Version defines manboster's application version.
const Version = "0.0.0"

// V indicates config's version, now is 0
const V = 0
