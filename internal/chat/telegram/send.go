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

	recp, err := recipientParser(msg.ChatID)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Telegram Provider] Error parsing recipient: %q", err))
		return err
	}

	opts := &telebot.SendOptions{}

	if msg.MessageID != "" {
		var replyID int
		_, err = fmt.Sscanf(msg.MessageID, "%d", &replyID)
		opts.ReplyTo = &telebot.Message{ID: replyID}
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Telegram Provider] Error getting reply id: %q", err))
			return err
		}
	}

	_, err = s.tgInstance.Send(recp, msg.Text.Text, opts)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Telegram Provider] Error sending message: %s", err))
		return err
	}
	color.Green(fmt.Sprintf("[Manboster Telegram Provider] Finally successfully sending message"))
	return nil
}
