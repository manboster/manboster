package telegram

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/go-viper/mapstructure/v2"
	"github.com/manboster/manboster/internal/chat"
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

	// Handler for Message Resp calling.
	s.tgInstance.Handle(telebot.OnText, func(c telebot.Context) error {
		msg := &chat.Message{
			Text:        c.Text(),
			MessageType: chat.MessageText,
			MessageID:   fmt.Sprintf("%d", c.Message().ID),
			Username:    c.Sender().FirstName + " " + c.Sender().LastName,
			UserID:      fmt.Sprintf("%d", c.Sender().ID),
			ChatID:      fmt.Sprintf("%d", c.Chat().ID),
			Provider:    "telegram",
		}

		typingCtx, cancelTyping := context.WithCancel(ctx)
		defer cancelTyping()
		go s.Type(telebot.ChatID(c.Chat().ID), typingCtx)

		// call onMsg on index
		onMsg(msg)
		return nil
	})

	color.Blue("Starting the telegram bot...")
	s.tgInstance.Start()
	return nil
}

func (s *Service) Notify(chatID string, action chat.ActionType) error {
	recipient, err := recipientParser(chatID)
	if err != nil {
		return err
	}
	return s.tgInstance.Notify(recipient, telebot.ChatAction(action))
}

func (s *Service) Type(chatId telebot.ChatID, ctx context.Context) {
	// 立即发一次
	_ = s.tgInstance.Notify(chatId, telebot.ChatAction(chat.ActionTyping))

	ticker := time.NewTicker(4 * time.Second) // 每 4 秒重发一次（略小于 5 秒过期时间）
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			_ = s.tgInstance.Notify(chatId, telebot.ChatAction(chat.ActionTyping))
		}
	}
}
