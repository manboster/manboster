package hook

import (
	"context"

	"github.com/manboster/manboster/spec/chat"
)

type EngineBeforeToolCallHookProvider struct {
	PolyfillFunc func(ctx context.Context, msg *chat.Message) (*chat.Message, error)
}

type EngineBeforeCompactHookProvider struct {
	PolyfillFunc func(ctx context.Context, before string) error
}

type EngineAfterCompactHookProvider struct {
	PolyfillFunc func(ctx context.Context, before string, after string) error
}

type EngineBeforeBuildSystemPromptHookProvider struct {
	PolyfillFunc func(ctx context.Context, before string) (string, error)
}
