package browser

import "github.com/manboster/manboster/spec/config"

type Config struct {
	Mode string `json:"mode" yaml:"mode" mapstructure:"mode" manboconfig:"name:browser mode;default:headless;desc:Browser launch mode. Use headless for servers; use headful only when a desktop UI is available.;desc_id:config.tool.browser.mode_desc" enum:"headful,headless"`
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
