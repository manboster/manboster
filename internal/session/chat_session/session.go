package chat_session

import (
	"context"
	"sync"

	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

// Session gives, stores and writes users session storages.
type Session struct {
	Events      []llm.Event
	Souls       []string
	Provider    string
	Model       string
	Command     chat.CommandType // if command is triggered, it would not be empty
	CommandStep int8             // the current step command is executing
	Active      bool
	Cancel      context.CancelFunc
	SessCancel  context.CancelFunc
	InputMsg    *chat.Message
	InputMsgID  string
	Ch          chan *chat.Message
}

type Manager struct {
	Lock     sync.RWMutex
	Sessions map[string]Session
}
