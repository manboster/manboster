package plugin

type EngineOperation struct {
	Type                 EngineOperationType              `json:"type"`
	InputResponsePayload *EngineOperationGetInputResponse `json:"inputResponsePayload,omitempty"`
	InputRequestPayload  *EngineOperationGetInputRequest  `json:"inputRequestPayload,omitempty"`
}

type EngineOperationGetInputRequest struct {
	Prompt string `json:"prompt"`
	Safety bool   `json:"safety"` // if safety is true, what the user input will be deleted later.
}

type EngineOperationGetInputResponse struct {
	Value string `json:"value"`
}

type EngineOperationType string

const (
	EngineOperationTypeNone EngineOperationType = ""
	EngineOperationGetInput EngineOperationType = "get_input"
)
