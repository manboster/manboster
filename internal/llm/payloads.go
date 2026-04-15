package llm

type MessageTextPayload struct {
	Text string `json:"text"`
}

type MessageToolRequestPayload struct {
	ID       string `json:"id"`
	ToolName string `json:"tool_name"` // the name you want to call
	ToolArgs any    `json:"tool_args"` // the args you want to call
}

type MessageToolResponsePayload struct {
	ID       string `json:"id"`
	ToolName string `json:"tool_name"`
	Result   string `json:"result"`
}
