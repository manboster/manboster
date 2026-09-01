package agent

import (
	"github.com/manboster/manboster/spec/llm"
)

// Chat chats with an agent, incoming marks a message for it
func (a *Agent) Chat(incoming llm.Message) (llm.Event, error) {
	return llm.Event{}, nil
}
