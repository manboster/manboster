package gateway

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
)

// SendSelect sends selection data with stateful maps and returns selection callback.
func (s *Service) SendSelect(ctx context.Context, instance chat.Provider, msg *chat.Message) (*chat.Message, error) {
	if msg.MessageType&chat.MessageSelection == 0 || msg.MessageType&chat.MessageTextAndImage == 0 {
		return nil, ErrInvalidMessageType
	}

	name := "sendselect_" + instance.Name() + "_" + msg.ChatID + "_" + strconv.FormatInt(time.Now().Unix(), 10)
	err := withRetry(ctx, name, 3, func(ctx context.Context) error {
		return nil
	})
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Gateway] failed to get selection from %q: %q", instance.DisplayName(), err))
		return nil, err
	}
	return nil, nil
}
