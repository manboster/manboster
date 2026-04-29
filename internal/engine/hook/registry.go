package hook

type Registry struct {
}

type Provider interface {
	HookType() EngineHookType
	PolyfillFunc(args ...interface{}) any
}

type EngineHookType string

const (
	EngineBeforeToolCall EngineHookType = "engine_before_tool_call"
)
