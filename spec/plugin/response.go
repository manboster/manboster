package plugin

import (
	"github.com/manboster/manboster/spec/chat"
)

type RunResponse struct {
	Response string // This response is for LLM.

	Hangup     bool               // is this hang up or not?
	Operations []EngineOperations // Operations wants to run
}

type EngineOperations struct {
	Type         EngineOperationType
	InputPayload *chat.Message
}

type EngineOperationType int

const (
	EngineOperationTypeNone EngineOperationType = iota
	EngineOperationGetInput
)
