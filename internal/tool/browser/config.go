package browser

import "github.com/manboster/manboster/spec/config"

type Config struct {
	Mode string `json:"mode" yaml:"mode" mapstructure:"mode" manboconfig:"name:browser mode;default:headless;desc:The mode of your browser, headless means this browser won't appear on your UI and you could not see it. Headful means you can see it on your Desktop and control it.\nBut if you deploy Manboster in VPS or desktop less environment, please do not use 'headful' as it's unavailable on these environment." enum:"headful,headless"`
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
