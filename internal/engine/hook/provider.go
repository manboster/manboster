package hook

import (
	"context"

	"github.com/manboster/manboster/spec/chat"
)

type Provider interface {
	HookType() EngineHookType
	PolyfillProvider() interface{}
}

type EngineHookType string

const (
	EngineBeforeToolCall EngineHookType = "engine_before_tool_call"
	EngineBeforeCompact  EngineHookType = "engine_before_compact"
	EngineAfterCompact   EngineHookType = "engine_after_compact"
)

type EngineBeforeToolCallHookProvider interface {
	PolyfillFunc(ctx context.Context, msg *chat.Message) (*chat.Message, error)
}

type EngineBeforeCompactHookProvider interface {
	PolyfillFunc(ctx context.Context, before string) error
}

type EngineAfterCompactHookProvider interface {
	PolyfillFunc(ctx context.Context, before string, after string) error
}
