package llm

// Message defines the chats between user and LLM.
type Message struct {
	Role RoleType    // Required. Who send it?
	Type MessageType // Required. What's the type of message?

	Text string // Optional. Required when MessageType = MessageText send text

	ToolName string // Optional. Required when MessageType = MessageToolCallRequest or MessageToolCallResponse , the name you want to call.
	ToolArgs any    // Optional. Required when MessageType = MessageToolCallRequest or MessageToolCallResponse, the params you want to call, or returns.
}

// MessageType defines Message's type.
type MessageType int16

const (
	MessageUnknown          MessageType = 0
	MessageText             MessageType = 1
	MessageToolCallRequest  MessageType = 255
	MessageToolCallResponse MessageType = 256
)

// RoleType is an enum defines role used by oai-compatible APIs.
type RoleType string

const (
	RoleSystem    RoleType = "system"
	RoleUser      RoleType = "user"
	RoleAssistant RoleType = "assistant"
	RoleToolCall  RoleType = "tool"
)
