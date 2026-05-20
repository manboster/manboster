package gguf

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/cli"
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

func (c *Config) Setup(ctx context.Context, p cli.Provider) error {
	confirm, err := p.Prompt("Do you want to enter model details manually?\n\nIf you don't know what's this, please select 'No'.", "Do you want to enter model details manually?", "Yes", "No")
	if err != nil {
		return err
	}

	if confirm {
		ggufURL, err := p.Input("GGUF File URL", "Please enter a valid URL to download GGUF.", "", false, func(input string) error {
			_, err := url.ParseRequestURI(input)
			if err != nil {
				return err
			}
			t := strings.HasSuffix(input, ".gguf")
			if t == false {
				return fmt.Errorf("gguf_url does not end with .gguf")
			}
			return nil
		})
		if err != nil {
			return err
		}

		sha256, err := p.Input("SHA256", "Please enter a valid SHA256 hash to download.\nIf you don't know what's this, please leave it as is.", "", false, func(input string) error { return nil })
		if err != nil {
			return err
		}

		c.GGUFurl = fmt.Sprintf("%s", ggufURL)
		c.GGUFsha256 = fmt.Sprintf("%s", sha256)

		return p.Alert("Manboster Configuration Wizard", "Successfully wrote hachimi config!")
	}

	err = p.Alert("Manboster Configuration Wizard", "Automatically set model with 'Qwen3-Guard Gen 0.6B'.")
	if err != nil {
		return err
	}
	c.ModelType = ModelSafeguard
	c.GGUFurl = models[0].Groups[0].Quants[3].URL
	c.GGUFsha256 = models[0].Groups[0].Quants[3].Sha256
	return nil
}
