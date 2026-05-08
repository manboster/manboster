package gatekeeper

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

func (s *Service) HachimiHandler(ctx context.Context, instance chat.Provider, msg *chat.Message, toolProvider tool.Provider, req llm.MessageToolCallRequestPayload, sid string) (bool, error) {
	if !*s.hachimiLoaded || s.hachimiProvider == nil {
		color.Yellow("[Manboster Gatekeeper] Hachimi is not loaded!")
		return true, nil
	}
	resp, err := s.hachimiProvider.Chat(ctx, "The model wants to run `rm -rf`.")
	if err != nil {
		return false, err
	}

	fmt.Printf("%+v\n", resp)

	return true, nil
}
