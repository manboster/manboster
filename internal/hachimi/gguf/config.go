package gguf

import (
	"context"
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/config"
)

type Config struct {
	GGUFurl       string       `json:"gguf_url" yaml:"gguf_url" mapstructure:"gguf_url" manboconfig:"skip"`
	GGUFsha256    string       `json:"gguf_sha256" yaml:"gguf_sha256" mapstructure:"gguf_sha256" manboconfig:"skip"`
	ModelType     ModelType    `json:"model_type" yaml:"model_type" mapstructure:"model_type" manboconfig:"required;enum;name:Model Type;desc:If you are using safeguard models, please select 'safeguard', otherwise, please leave it as is.;default:llm" enum:"safeguard,llm"`
	ContextLength ModelCtxType `json:"context_length" yaml:"context_length" mapstructure:"context_length" manboconfig:"required;enum;name:Model context length;desc:This value means how long context your hachimi can process, if your available RAM is low, please choose smaller one. The model's context is larger, the message can send is longer. If you don't know what's this, please leave it as is.;default:medium" enum:"low,medium,high,x-high"`
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

func (c *Config) Validate() error {
	if c.GGUFurl == "" {
		return fmt.Errorf("gguf_url is required")
	}
	if c.ContextLength == "" {
		c.ContextLength = ModelCtxMedium
		color.Yellow("[Manboster Hachimi Provider] Context length is not found, loading it as default value medium")
	}
	if c.ModelType == "" {
		return fmt.Errorf("model_type is required")
	}
	return nil
}

func (c *Config) Setup(ctx context.Context) error {
	sel := false
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().Title("Do you want to enter model details manually?").Description("If you don't know what's this, please select 'No'.").Negative("No").Affirmative("Yes").Value(&sel),
		)).Run()
	if err != nil {
		return err
	}

	if sel {
		// TODO: add user defined sha256 and gguf link
		color.Blue("Work in Progress...")
		return nil
	}

	color.Blue("Set your model to qwen3 guard 0.6B, other models is work in progress...")
	// TODO: setup
	c.ModelType = ModelSafeguard
	c.GGUFurl = models[0].Groups[0].Quants[2].URL
	c.GGUFsha256 = models[0].Groups[0].Quants[2].Sha256
	return nil
}
