package llm

// Message defines the chats between user and LLM.
type Message struct {
	Role RoleType    // Required. Who send it?
	Type MessageType // Required. What's the type of message?

	Parts []MessageParts // Optional. Required when MessageType = MessageText || MessageImage || MessageFile send text

	ToolCallRequest  []MessageToolCallRequestPayload  // Optional. Required when MessageType = MessageToolCallRequest
	ToolCallResponse []MessageToolCallResponsePayload // Optional. Required when MessageType = MessageToolCallResponse

	Thinking *MessageThinkingPayload // Optional. Required when Message & MessageThinking != 0. Model's Thinking Content.
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
	MessageThinking
)

// RoleType is an enum defines role used by oai-compatible APIs.
type RoleType string

const (
	RoleSystem    RoleType = "system"
	RoleUser      RoleType = "user"
	RoleAssistant RoleType = "assistant"
	RoleToolCall  RoleType = "tool"
)
