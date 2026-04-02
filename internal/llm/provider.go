package llm

import "context"

// Provider defines which functions should LLM provides give.
type Provider interface {
	Chat(ctx context.Context, messages []Message) (*Message, error) // now one.
	Init(ctx context.Context, config any) error
	Name() string
	Model() string
	New() Provider
}

// Message defines the chats between user and LLM.
type Message struct {
	Role RoleType // Who send it?
	Type MessageType
	Text string // send text
}

// MessageType defines Message's type.
type MessageType int16

const (
	MessageTypeUnknown          MessageType = 0
	MessageTypeText             MessageType = 1
	MessageTypeToolCallRequest  MessageType = 255
	MessageTypeToolCallResponse MessageType = 256
)

// RoleType is an enum defines role used by oai-compatible APIs.
type RoleType string

const (
	RoleTypeSystem    RoleType = "system"
	RoleTypeUser      RoleType = "user"
	RoleTypeAssistant RoleType = "assistant"
	RoleTypeToolCall  RoleType = "tool"
)
