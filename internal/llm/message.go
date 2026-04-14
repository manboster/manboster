package llm

// Message defines the chats between user and LLM.
type Message struct {
	Role RoleType    // Required. Who send it?
	Type MessageType // Required. What's the type of message?

	Text *MessageTextPayload // Optional. Required when MessageType = MessageText send text

	Tool *MessageToolPayload // Optional. Required when MessageType = MessageToolCallRequest or MessageToolCallResponse
}

// MessageType defines Message's type.
type MessageType int16

const (
	MessageUnknown MessageType = 1 << iota
	MessageText
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
