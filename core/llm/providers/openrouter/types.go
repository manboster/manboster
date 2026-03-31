package openrouter

import "github.com/charmbracelet/huh"

// Config contains what you should enter in application configuration.
type Config struct {
	ApiKey string `yaml:"api_key"` // your openrouter system's apikey
	// BaseURL string `yaml:"base_url"` // this is fixed so you don't need to enter it.
	Model string `yaml:"model"` // your wanted model like anthropic/claude-sonnet-4.5
}

// models defines some recently used popular models so you don't search it in openrouter.
var models = []string{
	"openrouter/auto",
	"openrouter/free",
	"google/gemini-3.1-pro-preview",
	"google/gemini-3-flash-preview",
	"anthropic/claude-sonnet-4.6",
	"anthropic/claude-opus-4.6",
	"anthropic/claude-sonnet-4.5",
	"anthropic/claude-opus-4.5",
	"anthropic/claude-haiku-4.5",
	"openai/gpt-5.4-nano",
	"openai/gpt-5.4-mini",
	"openai/gpt-5.4-pro",
	"openai/gpt-5.4",
	"openai/gpt-5.3-chat",
	"openai/gpt-5.3-codex",
	"openai/gpt-oss-20b:free",
	"openai/gpt-oss-120b:free",
	"openai/gpt-oss-20b",
	"openai/gpt-oss-120b",
	"moonshotai/kimi-k2.5",
	"deepseek/deepseek-v3.2",
	"stepfun/step-3.5-flash:free",
	"stepfun/step-3.5-flash",
	"minimax/minimax-m2.7",
	"minimax/minimax-m2.5:free",
	"minimax/minimax-m2.5",
	"z-ai/glm-5-turbo",
	"z-ai/glm-5",
	"x-ai/grok-4",
	"x-ai/grok-4.1-fast",
	"qwen/qwen3.5-397b-a17b",
	"qwen/qwen3.5-flash-02-23",
	"xiaomi/mimo-v2-pro",
	"xiaomi/mimo-v2-flash",
}

// ToHuhGroup enables configuration go ahead.
func (c *Config) ToHuhGroup() []*huh.Group {
	var modelOptions []huh.Option[string]
	for _, m := range models {
		modelOptions = append(modelOptions, huh.NewOption(m, m))
	}
	modelOptions = append(modelOptions, huh.NewOption("Other Model", "CustomModel"))

	return []*huh.Group{
		huh.NewGroup(
			huh.NewInput().Title("Your OpenRouter API Key").Description("Your OpenRouter API Key.\nIf you don't have one, please open https://openrouter.ai/workspaces/default/keys to create one.").Value(&c.ApiKey),
			huh.NewSelect[string]().Title("OpenRouter Models").Description("Select the model you want to use as Manboster's brain.").Options(
				modelOptions...,
			).Value(&c.Model),
		),
	}
}

// GetConfig returns its own struct.
func (c *Config) GetConfig() any {
	return c
}

// VerifyAndConvert ensures configuration is valid.
func (c *Config) VerifyAndConvert() error {
	// If you choose Custom Model, you should specify it.
	if c.Model == "CustomModel" {
		err := huh.NewForm(huh.NewGroup(huh.NewInput().Title("Your Model Name").Description("Please specify the model name. You can copy it by clicking the clipboard icon on OpenRouter's model page.").Value(&c.Model))).Run()
		if err != nil {
			return err
		}
	}
	return nil
}
