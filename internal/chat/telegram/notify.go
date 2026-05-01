package telegram

import (
	"context"
	"fmt"
	"strconv"

	"github.com/manboster/manboster/spec/chat"
	"gopkg.in/telebot.v3"
)

func (s *Service) Notify(ctx context.Context, msg *chat.Message, action chat.ActionType) error {
	isEnable := false
	isClean := false
	switch s.cfg.ReactionNotifyStatus {
	case "disabled":
	case "enabled":
		isEnable = true
	case "clean":
		isClean = true
		isEnable = true
	default:
	}

	// mark it reaction
	recipient, err := recipientParser(msg.ChatID)
	if err != nil {
		return err
	}

	msgId := 0
	chatId := int64(0)
	msgId, err = strconv.Atoi(msg.MessageID)
	if err != nil {
		msgId = 0
	}

	chatId, err = strconv.ParseInt(msg.ChatID, 10, 64)
	if err != nil {
		return err
	}

	switch action {
	case chat.ActionPending:
		typingCtx, cancelTyping := context.WithCancel(ctx)
		notifierWrite(chatId, msgId, cancelTyping)
		go s.Type(typingCtx, telebot.ChatID(chatId))

		if isEnable {
			return s.tgInstance.React(recipient, &telebot.Message{
				ID: msgId,
				Chat: &telebot.Chat{
					ID: chatId,
				},
			}, telebot.ReactionOptions{
				Reactions: []telebot.Reaction{
					{
						Type:  "emoji",
						Emoji: "✍️",
					},
				},
			})
		}
		return nil

	case chat.ActionSuccess:
		chatId, msgId = notifierCancel(msg.ChatID)

		if isClean {
			return s.tgInstance.React(recipient, &telebot.Message{
				ID: msgId,
				Chat: &telebot.Chat{
					ID: chatId,
				},
			}, telebot.ReactionOptions{
				Reactions: []telebot.Reaction{},
			})
		}
		return nil

	case chat.ActionError:
		notifierCancel(msg.ChatID)
		return nil
	default:
		return fmt.Errorf("invalid action type: %v", action)
	}
}
