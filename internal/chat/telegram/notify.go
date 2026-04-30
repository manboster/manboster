package telegram

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/spec/chat"
	"gopkg.in/telebot.v3"
)

func (s *Service) Notify(ctx context.Context, msg *chat.Message, action chat.ActionType) error {
	switch s.cfg.ReactionNotifyStatus {
	case "disabled":
		return nil
	case "enabled", "clean":
	default:
		return nil
	}
	switch action {
	case chat.ActionPending:
		// mark it reaction
		recipient, err := recipientParser(msg.ChatID)
		if err != nil {
			return err
		}

		chatId := int64(0)
		_, err = fmt.Sscanf(msg.ChatID, "%d", &chatId)
		if err != nil {
			return err
		}

		msgId := 0
		_, err = fmt.Sscanf(msg.MessageID, "%d", &msgId)

		if err != nil {
			return err
		}

		typingCtx, cancelTyping := context.WithCancel(ctx)
		notifierWrite(chatId, msgId, cancelTyping)
		go s.Type(typingCtx, telebot.ChatID(chatId))

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
	case chat.ActionSuccess:
		recipient, err := recipientParser(msg.ChatID)
		if err != nil {
			return err
		}

		cid, mid := notifierCancel(msg.ChatID)
		return s.tgInstance.React(recipient, &telebot.Message{
			ID: mid,
			Chat: &telebot.Chat{
				ID: cid,
			},
		}, telebot.ReactionOptions{})
	case chat.ActionError:
		notifierCancel(msg.ChatID)
		return nil
	default:
		return fmt.Errorf("invalid action type: %v", action)
	}
}
