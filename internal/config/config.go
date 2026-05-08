package config

// Config makes configurations to this whole application and store it in yml format.
type Config struct {
	Version int16          `yaml:"version" mapstructure:"version" json:"version"`  // version data, used for migration.
	Chats   []ChatConfig   `yaml:"chats" mapstructure:"chats" json:"chats"`        // chat configurations
	LLMs    []LLMConfig    `yaml:"llms" mapstructure:"llms" json:"llms"`           // LLM configurations
	Tools   []ToolConfig   `yaml:"tools" mapstructure:"tools" json:"tools"`        // Tool configurations
	App     AppConfig      `yaml:"app" mapstructure:"app" json:"app"`              // APP specific configurations
	Hachimi HachimiConfigs `yaml:"hachimi" mapstructure:"hachimi" json:"hachimi" ` // Hachimi configurations
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
	DBPath             string `yaml:"dbpath" mapstructure:"dbpath" json:"dbpath"`                                           // SQLite's path
	DefaultLLMProvider string `yaml:"default_llm_provider" mapstructure:"default_llm_provider" json:"default_llm_provider"` // default llm provider
	DefaultLLMModel    string `yaml:"default_llm_model" mapstructure:"default_llm_model" json:"default_llm_model"`          // default model
}

// ToolConfig is configuration available to tool(local), plugin(external).
type ToolConfig struct {
	Name          string `yaml:"name" mapstructure:"name" json:"name"`
	Configuration any    `yaml:"configuration" mapstructure:"configuration" json:"configuration,omitempty"`
}

// HachimiConfig is configuration available to hachimi providers
type HachimiConfig struct {
	Provider      string `yaml:"name" mapstructure:"name" json:"name"`
	Configuration any    `yaml:"configuration" mapstructure:"configuration" json:"configuration"`
}

// HachimiConfigs wraps HachimiConfig
type HachimiConfigs struct {
	Enabled  bool            `yaml:"enabled" mapstructure:"enabled" json:"enabled"`
	Provider string          `yaml:"provider" mapstructure:"provider" json:"provider"` // default provider
	Hachimi  []HachimiConfig `yaml:"hachimi" mapstructure:"hachimi" json:"hachimi"`
}
