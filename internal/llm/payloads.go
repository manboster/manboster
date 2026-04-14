package llm

type MessageTextPayload struct {
	Text string `json:"text"`
}

type MessageToolPayload struct {
	ToolName string `json:"tool_name"` // the name you want to call
	ToolArgs any    `json:"tool_args"` // the args you want to call
}
