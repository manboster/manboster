package llm

type MessageTextPayload struct {
	Text string `json:"text"`
}

type MessageToolCallRequestPayload struct {
	ID       string `json:"id"`
	ToolName string `json:"tool_name"` // the name you want to call
	ToolArgs any    `json:"tool_args"` // the args you want to call
}

type MessageToolCallResponsePayload struct {
	ID       string `json:"id"`
	ToolName string `json:"tool_name"`
	Result   string `json:"result"`
}

type MessageThinkingPayload struct {
	Thinking string `json:"thinking"`
}

type MessageImagePayload struct {
	Content string `json:"content"`
}
