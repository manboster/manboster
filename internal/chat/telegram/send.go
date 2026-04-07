package telegram

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"gopkg.in/telebot.v3"
)

// SendMessage sends a message to user.
func (s *Service) SendMessage(ctx context.Context, msg *chat.Message) error {
	// temporarily solution
	s.sendMutex.Lock()
	defer s.sendMutex.Unlock()

	i := 1
	for i <= 5 {
		recp, err := recipientParser(msg.ChatID)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Telegram Provider]Try %d times, error parsing recipient: %s", i, err))
			i += 1
			continue
		}

		opts := &telebot.SendOptions{}

		if msg.MessageID != "" {
			var replyID int
			_, err := fmt.Sscanf(msg.MessageID, "%d", &replyID)
			opts.ReplyTo = &telebot.Message{ID: replyID}
			if err != nil {
				color.Red(fmt.Sprintf("[Manboster Telegram Provider]Try %d times, error getting reply id: %s", i, err))
				i += 1
				continue
			}
		}

		_, err = s.tgInstance.Send(recp, msg.Text.Text, opts)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Telegram Provider]Try %d times, error sending message: %s", i, err))
			i += 1
			continue
		} else {
			color.Green(fmt.Sprintf("[Manboster Telegram Provider]Try %d times, finally successfully sending message", i))
			return nil
		}
	}
	color.Red(fmt.Sprintf("[Manboster Telegram Provider]Try %d times, can not send telegram message!", i))
	return ErrSendFailed
}
