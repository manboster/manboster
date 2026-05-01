package telegram

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
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

	opts := &telebot.SendOptions{
		ParseMode: telebot.ModeHTML,
	}

	if msg.Reply != nil {
		var replyID int
		_, err = fmt.Sscanf(msg.Reply.MessageID, "%d", &replyID)
		opts.ReplyTo = &telebot.Message{ID: replyID}
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Telegram Provider] Error getting reply id: %q", err))
			return err
		}
	}

	isToolCall := strings.HasPrefix(msg.Text.Text, "Model called")
	text := s.Converter(msg.Text.Text, msg.MessageType&chat.MessageThinkingText != 0, isToolCall)
	limit := 4000
	// check length of the text and slice it
	if utf8.RuneCountInString(text) < limit {
		m, sendErr := s.tgInstance.Send(recp, text, opts)
		if sendErr != nil {
			color.Yellow(fmt.Sprintf("[Manboster Telegram Provider] Error sending message: %q", err))
			m, err = s.tgInstance.Send(recp, msg.Text.Text, &telebot.SendOptions{
				ParseMode:             telebot.ModeDefault,
				DisableWebPagePreview: true,
				DisableNotification:   isToolCall,
			})
		}
		if m != nil {
			msg.MessageID = fmt.Sprintf("%d", m.ID)
		}
		color.Green(fmt.Sprintf("[Manboster Telegram Provider] Finally successfully sending message"))
	} else {
		caption := "We are sorry but the message is too long to send, please open message.txt above to read it.\n"
		if msg.MessageType == chat.MessageThinkingText {
			caption += "This is the thinking process of this model."
		} else {
			caption += "This is the response of this model."
		}
		m, sendErr := s.tgInstance.Send(recp, &telebot.Document{
			Caption:  caption,
			FileName: "message.txt",
			File:     telebot.FromReader(strings.NewReader(msg.Text.Text)),
		}, opts)
		if sendErr != nil {
			color.Red(fmt.Sprintf("[Manboster Telegram Provider] Error sending message: %s", err))
			return err
		}
		if m != nil {
			msg.MessageID = fmt.Sprintf("%d", m.ID)
		}
		color.Yellow(fmt.Sprintf("[Manboster Telegram Provider] Message is too long! We can't send it via text, however, we sent it via file."))
	}

	return nil
}
