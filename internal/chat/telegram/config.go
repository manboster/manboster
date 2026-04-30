package telegram

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/config"
)

// Config configures their Telegram bot.
type Config struct {
	BotToken             string `yaml:"bot_token" json:"bot_token" mapstructure:"bot_token" manboconfig:"required;secret;name:Telegram Bot Token;desc:Your Telegram Bot Token, if you don't have one，you can get it from @BotFather."` // Telegram requires your bot token to authenticate their server.
	CollapseMsgLength    int16  `yaml:"collapse_msg_length" json:"collapse_msg_length" mapstructure:"collapse_msg_length" manboconfig:"required;secret;name:Collapse Message Length;desc:The length to collapse message with an expandable quote when telegram's message is longer than of which.;default:500"`
	ReactionNotifyStatus string `yaml:"reaction_notify_status" json:"reaction_notify_status" mapstructure:"reaction_notify_status" manboconfig:"required;name:Reaction Notify Status;desc:The preferences of select a notification.\nThe default value is enabled.\nEnable means emoji reaction on your message when processing and the clean one will wrap it up from the original message after the message is successfully handled. Disable means no emoji reaction on your message when processing.;default:enabled" enum:"disabled,enabled,clean"`
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
	if c.CollapseMsgLength == 0 {
		c.CollapseMsgLength = 500
		color.Yellow("[Manboster Telegram Provider] could not read collapse message length, setting it to default value 500.")
	}
	if c.CollapseMsgLength > 3500 {
		c.CollapseMsgLength = 3500
		color.Yellow("[Manboster Telegram Provider] the length is too long for configuration! Setting it to maximum value 3500")
	}
	if c.ReactionNotifyStatus != "enabled" && c.ReactionNotifyStatus != "clean" && c.ReactionNotifyStatus != "disabled" {
		c.ReactionNotifyStatus = "enabled"
		color.Yellow("[Manboster Telegram Provider] could not read reaction notify status, setting it to default value 'enabled'.")
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
