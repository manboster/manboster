package runner

import (
	"github.com/manboster/manboster/internal/engine/gateway"
	"github.com/manboster/manboster/internal/engine/processor"
	"github.com/manboster/manboster/spec/chat"
)

type Runner struct {
	InputCh       chan MsgData
	processor     *processor.Service
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

func NewRunner(processorService *processor.Service, providerMap map[string]chat.Provider) *Runner {
	return &Runner{
		InputCh:       InputCh,
		processor:     processorService,
		chatProviders: providerMap,
	}
}

var Instance *Runner // Activate it in engine.load
var InputCh = make(chan MsgData, 32)
