package gguf

import "github.com/manboster/manboster/spec/config"

type Config struct {
}

func (c *Config) Name() string {
	return "hachimi-gguf"
}

func (c *Config) DisplayName() string {
	return "hachimi gguf runtime"
}

func (c *Config) Args() *config.Args {
	return config.ArgsFromStruct(Config{})
}

func (c *Config) GetConfig() any {
	return c
}

func (c *Config) Validate() {

}

func (c *Config) Setup() {

}
