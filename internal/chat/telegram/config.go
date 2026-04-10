package telegram

import (
	"context"
	"errors"
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/manboster/manboster/internal/util"
)

// Config configures their Telegram bot.
type Config struct {
	BotToken string `yaml:"bot_token" json:"bot_token" mapstructure:"bot_token"` // Telegram requires your bot token to authenticate their server.
}

// ToHuhGroup enables configuration go ahead.
func (c *Config) ToHuhGroup() []*huh.Group {
	return []*huh.Group{
		huh.NewGroup(
			huh.NewInput().Title("Telegram Bot Token").Description("Your Telegram Bot's token.\nIf you don't have any, please open your Telegram, search @BotFather and create one.").EchoMode(huh.EchoModePassword).Value(&c.BotToken)),
	}
}

// VerifyAndConvert helps to convert internal string ru to RootUser int64 type.
func (c *Config) VerifyAndConvert(ctx context.Context) error {
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
