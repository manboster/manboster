package telegram

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/go-viper/mapstructure/v2"
	"github.com/manboster/manboster/core/chat"
	"gopkg.in/telebot.v3"
)

// Start starts your Telegram Service
func (s *Service) Start(ctx context.Context, conf any, onMsg func(msg *chat.Message)) error {
	// get config
	var cfg Config
	err := mapstructure.Decode(conf, &cfg)
	if err != nil {
		return err
	}

	// validate
	if cfg.BotToken == "" {
		return ErrBotTokenRequired
	}

	// set the bot and start it.
	settings := telebot.Settings{
		Token:  cfg.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	// start the bot
	b, err := telebot.NewBot(settings)
	if err != nil {
		return err
	}
	s.tgInstance = b

	// ctx done cleaning
	go func() {
		_ = s.Stop(ctx)
	}()

	//
	s.tgInstance.Handle(telebot.OnText, func(c telebot.Context) error {
		msg := &chat.Message{
			Text:        c.Text(),
			MessageType: chat.MessageTypeText,
			MessageID:   fmt.Sprintf("%d", c.Message().ID),
			Username:    c.Sender().FirstName + " " + c.Sender().LastName,
			UserID:      fmt.Sprintf("%d", c.Sender().ID),
			ChatID:      fmt.Sprintf("%d", c.Chat().ID),
			Provider:    "telegram",
		}

		// 触发 main 函数传入的回调逻辑 (例如调用 LLM)
		onMsg(msg)
		return nil
	})

	color.Blue("Starting the telegram bot...")
	s.tgInstance.Start()
	return nil
}
