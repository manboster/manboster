package gguf

import "github.com/manboster/manboster/spec/config"

type Config struct {
	GGUFurl    string `json:"gguf_url" yaml:"gguf_url" mapstructure:"gguf_url" manboconfig:"skip"`
	GGUFsha256 string `json:"gguf_sha256" yaml:"gguf_sha256" mapstructure:"gguf_sha256" manboconfig:"skip"`
}

type ModelType string

const (
	ModelSafeguard ModelType = "safeguard"
	ModelLLM       ModelType = "llm"
)

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
	// TODO: setup
	c.GGUFurl = models[0].Groups[0].Quants[2].URL
	c.GGUFsha256 = models[0].Groups[0].Quants[2].Sha256
}
