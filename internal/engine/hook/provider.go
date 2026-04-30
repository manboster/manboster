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
)

type EngineBeforeToolCallHookProvider interface {
	HookType() EngineHookType
	PolyfillFunc(ctx context.Context, msg *chat.Message) (*chat.Message, error)
}
