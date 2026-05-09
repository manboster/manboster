package runner

import (
	"context"

	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

type requiredInterface interface {
	MessageHandler(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error
}

type Runner struct {
	InputCh chan MsgData
	engine  requiredInterface
}

type MsgData struct {
	SessionID string
	Instance  string
	ChatMsg   *chat.Message
	LLMMsg    *llm.Message
}

func NewRunner(e requiredInterface) *Runner {
	return &Runner{
		InputCh: make(chan MsgData, 16),
		engine:  e,
	}
}

var Instance *Runner // Activate it in engine.load
