package llm

import "context"

// Provider defines which functions should LLM provides give.
type Provider interface {
	Chat(ctx context.Context, messages []Message) (*Message, error)              // now one.
	ChatStream(ctx context.Context, messages []Message) (<-chan *Message, error) // WIP: New Streaming chat
	Init(ctx context.Context, config any) error
	Name() string
	Model() string
	New() Provider
}

// Message defines the chats between user and LLM.
type Message struct {
	Role RoleType    // Required. Who send it?
	Type MessageType // Required. What's the type of message?

	Text string // Optional. Required when MessageType = MessageTypeText send text

	ToolName string // Optional. Required when MessageType = MessageTypeToolCallRequest or MessageTypeToolCallResponse , the name you want to call.
	ToolArgs any    // Optional. Required when MessageType = MessageTypeToolCallRequest or MessageTypeToolCallResponse, the params you want to call, or returns.
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
