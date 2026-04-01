package telegram

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/manboster/manboster/core/util"
)

// Config configures their Telegram bot.
type Config struct {
	BotToken string `yaml:"bot_token" json:"bot_token" mapstructure:"bot_token"` // Telegram requires your bot token to authenticate their server.
	RootUser int64  `yaml:"root_user" json:"root_user" mapstructure:"root_user"` // Telegram uid used to authenticate root access.
	ru       string // this is used to transfer data.
}

// ToHuhGroup enables configuration go ahead.
func (c *Config) ToHuhGroup() []*huh.Group {
	return []*huh.Group{
		huh.NewGroup(
			huh.NewInput().Title("Telegram Bot Token").Description("Your Telegram Bot's token.\nIf you don't have any, please open your Telegram, search @BotFather and create one.").EchoMode(huh.EchoModePassword).Value(&c.BotToken),
			huh.NewInput().Title("Your Telegram UID").Description("Your Telegram UID.\nIf you don't know what it is, you can just make it empty, send '/id' to your bot after the service started and run 'manboster root [your uid]' to authenticate instead.").Value(&c.ru),
		),
	}
}

// VerifyAndConvert helps to convert internal string ru to RootUser int64 type.
func (c *Config) VerifyAndConvert() error {
	if c.ru != "" {
		rUid, err := strconv.ParseInt(c.ru, 10, 64)
		if err != nil {
			return err
		}
		c.RootUser = rUid
	}
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
	return fmt.Sprintf("BotToken: %s, RootUser: %d", util.MaskSecret(c.BotToken), c.RootUser)
}
