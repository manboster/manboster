package cron

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/manboster/manboster/internal/engine/runner"
	"github.com/manboster/manboster/spec/chat"
)

func isDelayFormat(s string) bool {
	if !strings.HasPrefix(s, "+") {
		return false
	}
	raw := s[1:]
	if raw == "" {
		return false
	}
	switch raw[len(raw)-1] {
	case 's', 'm', 'h', 'd':
		_, err := strconv.Atoi(raw[:len(raw)-1])
		return err == nil
	default:
		return false
	}
}

func parseDelay(s string) (time.Duration, error) {
	if !strings.HasPrefix(s, "+") {
		return 0, fmt.Errorf("invalid delay format: %s", s)
	}
	raw := s[1:]
	if strings.HasSuffix(raw, "d") {
		daysStr := strings.TrimSuffix(raw, "d")
		days, err := strconv.Atoi(daysStr)
		if err != nil {
			return 0, fmt.Errorf("invalid delay format: %s", s)
		}
		return time.Duration(days) * 24 * time.Hour, nil
	}
	return time.ParseDuration(raw)
}

func buildMessageDataFromArgs(arg RunArgs, chatId string, chatProvider string, userId string) runner.MsgData {
	var msgData runner.MsgData
	msgData.ChatMsg = &chat.Message{
		Provider:    chatProvider,
		MessageType: chat.MessageText,
	}

	switch arg.MessageType {
	case MessageText:
		msgData.Type = runner.MsgText
	case MessagePrompt:
		msgData.Type = runner.MsgPrompt
	default:
		msgData.Type = runner.MsgText
	}

	switch arg.To {
	case ToThisChat:
		msgData.ChatMsg.ChatID = chatId
	case ToPM:
		msgData.ChatMsg.ChatID = userId
	default:
		msgData.ChatMsg.ChatID = chatId
	}

	switch arg.Ignore {
	case IgnoreNone:
		msgData.ChatMsg.MessageType |= chat.MessageFromCronIgnore
	case IgnoreHachimi:
		msgData.ChatMsg.MessageType |= chat.MessageFromCron
	default:
		msgData.ChatMsg.MessageType |= chat.MessageFromCron
	}

	msgData.ChatMsg.Text = &chat.TextPayload{
		Text: arg.Prompt,
	}

	return msgData
}
