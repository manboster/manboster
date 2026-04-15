package llm

// Message defines the chats between user and LLM.
type Message struct {
	Role RoleType    // Required. Who send it?
	Type MessageType // Required. What's the type of message?

	Parts []MessageParts // Optional. Required when MessageType = MessageText || MessageImage || MessageFile send text

	ToolRequest  *MessageToolRequestPayload  // Optional. Required when MessageType = MessageToolCallRequest
	ToolResponse *MessageToolResponsePayload // Optional. Required when MessageType = MessageToolCallResponse
}

// MessageType defines Message's type.
type MessageType int16

const (
	MessageText MessageType = 1 << iota
	MessageImage
	MessageAudio
	MessageVideo
	MessageFile
	MessageToolCallRequest
	MessageToolCallResponse
)

// RoleType is an enum defines role used by oai-compatible APIs.
type RoleType string

const (
	RoleSystem    RoleType = "system"
	RoleUser      RoleType = "user"
	RoleAssistant RoleType = "assistant"
	RoleToolCall  RoleType = "tool"
)
