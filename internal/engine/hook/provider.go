package hook

type EngineHookType string

const (
	EngineBeforeToolCall          EngineHookType = "engine_before_tool_call"
	EngineAfterToolCall           EngineHookType = "engine_after_tool_call"
	EngineBeforeChat              EngineHookType = "engine_before_chat"
	EngineAfterChat               EngineHookType = "engine_after_chat"
	EngineBeforeCompact           EngineHookType = "engine_before_compact"
	EngineAfterCompact            EngineHookType = "engine_after_compact"
	EngineBeforeBuildSystemPrompt EngineHookType = "engine_before_build_system_prompt"
)
