package file

import (
	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/config"
)

type Config struct {
	Mode string `json:"mode" yaml:"mode" mapstructure:"mode" manboconfig:"name:file write mode;default:readonly;desc:The file write mode, readonly means model couldn't write and delete anything in the workspace. Readwrite enables model to write or delete data in the workspace.'" enum:"readonly,readwrite"`
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
	if c.Mode != "readonly" && c.Mode != "readwrite" {
		c.Mode = "readonly"
		color.Yellow("[Manboster Tool Provider] Could not read mode, setting it to default value 'readonly'.")
	}
	return nil
}

func (c *Config) GetConfig() any {
	return c
}
