package gatekeeper

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
)

func (s *Service) RecallRunner(ctx context.Context, instance chat.Provider, msg *chat.Message, t time.Duration) error {
	timer := time.NewTimer(t)
	defer timer.Stop()

	forkedMsg := msg.Clone()
	forkedMsg.MessageID = msg.MessageID
	forkedMsg.ChatID = msg.ChatID
	forkedMsg.MessageType = chat.MessageUnknown

	for {
		select {
		case <-ctx.Done():
			err := instance.DeleteMessage(ctx, forkedMsg)
			if err != nil {
				color.Yellow(fmt.Sprintf("Failed to delete message: %v", err))
				return err
			}
			return ctx.Err()
		case <-timer.C:
			err := instance.DeleteMessage(ctx, forkedMsg)
			if err != nil {
				color.Yellow(fmt.Sprintf("Failed to delete message: %v", err))
				return err
			}
			return nil
		}
	}
}

func (s *Service) Recall(ctx context.Context, instance chat.Provider, rMsg *chat.Message) {
	err := s.RecallRunner(ctx, instance, rMsg, 5*time.Second)
	if err != nil {
		color.Yellow("[Manboster Gatekeeper] Failed to recall result")
	}
}
