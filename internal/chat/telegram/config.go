package telegram

import (
	"errors"
	"fmt"

	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/config"
)

// Config configures their Telegram bot.
type Config struct {
	BotToken string `yaml:"bot_token" json:"bot_token" mapstructure:"bot_token" manboconfig:"required,secret,desc:Your Telegram Bot Token"` // Telegram requires your bot token to authenticate their server.
}

// Args return args to write
func (c *Config) Args() *config.Args {
	return config.ArgsFromStruct(Config{})
}

// Validate validates configuration data
func (c *Config) Validate() error {
	if c.BotToken == "" {
		return errors.New("bot token is required")
	}
	return nil
}

// GetConfig returns itself directly to the app.
func (c *Config) GetConfig() any {
	return c
}

// String is used to print sth.
func (c *Config) String() string {
	return fmt.Sprintf("BotToken: %s", util.MaskSecret(c.BotToken))
}

func (c *Config) Name() string {
	return "telegram"
}

func (c *Config) DisplayName() string {
	return "telegram"
}
