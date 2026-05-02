package hook

import (
	"context"

	"github.com/manboster/manboster/spec/chat"
)

type EngineHookType string

const (
	EngineBeforeToolCall EngineHookType = "engine_before_tool_call"
	EngineBeforeCompact  EngineHookType = "engine_before_compact"
	EngineAfterCompact   EngineHookType = "engine_after_compact"
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
