package runner

import (
	"context"

	"github.com/manboster/manboster/internal/engine/gateway"
	"github.com/manboster/manboster/spec/chat"
)

type requiredInterface interface {
	HandleMessage(ctx context.Context, instance chat.Provider, msg *chat.Message) error
}

type Runner struct {
	InputCh       chan MsgData
	engine        requiredInterface
	gateway       *gateway.Service
	chatProviders map[string]chat.Provider
}

type MsgData struct {
	Type    MsgType
	ChatMsg *chat.Message
}

type MsgType string

const (
	MsgPrompt MsgType = "prompt"
	MsgText   MsgType = "text"
)

func NewRunner(e requiredInterface, providerMap map[string]chat.Provider) *Runner {
	return &Runner{
		InputCh:       make(chan MsgData, 16),
		engine:        e,
		chatProviders: providerMap,
	}
}

var Instance *Runner // Activate it in engine.load
