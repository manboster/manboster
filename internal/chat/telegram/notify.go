package telegram

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/chat"
	"gopkg.in/telebot.v3"
)

func (s *Service) Notify(ctx context.Context, msg *chat.Message, action chat.ActionType) error {
	switch action {
	case chat.ActionTyping:
		recipient, err := recipientParser(msg.ChatID)
		if err != nil {
			return err
		}
		return s.tgInstance.Notify(recipient, telebot.ChatAction(action))
	case chat.ActionPending:
		recipient, err := recipientParser(msg.ChatID)
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
		// TODO
		return nil
	case chat.ActionError:
		// TODO
		return nil
	default:
		return fmt.Errorf("invalid action type: %v", action)
	}
}
