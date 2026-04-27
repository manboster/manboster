package gateway

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
)

// SendSelect sends selection data with stateful maps and returns selection callback.
func (s *Service) SendSelect(ctx context.Context, instance chat.Provider, msg *chat.Message) (*chat.Message, error) {
	if msg.MessageType&chat.MessageSelection == 0 {
		return nil, ErrInvalidMessageType
	}

	sid := util.RandomString(8)
	s.selectionSessionManager.SetSelectMsg(sid, msg)
	defer s.selectionSessionManager.CleanSelect(sid)

	ch := make(chan *chat.Message, 1)
	s.selectionSessionManager.SetSelectionChan(sid, ch)
	name := "sendselect_" + instance.Name() + "_" + msg.ChatID + "_" + strconv.FormatInt(time.Now().Unix(), 10)
	err := withRetry(ctx, name, 3, func(ctx context.Context) error {
		timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Minute)
		defer cancel()

		return instance.Select(timeoutCtx, sid, msg)
	})

	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Gateway] failed to get selection from %q: %q", instance.DisplayName(), err))
		return nil, err
	}

	select {
	case respMsg := <-ch:
		return respMsg, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(5 * time.Minute):
		return nil, ErrTimeout
	}
}
