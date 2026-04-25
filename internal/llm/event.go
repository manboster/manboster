package llm

// Event defines payloads of responses
type Event struct {
	EventType EventType // Required.
	Model     string    // which model returned this.
	Provider  string    // Provider of this LLM model

	Message *Message // Optional. Required when EventType & EventUsage != 0
	Usage   *Usage   // Optional. Required when EventType & EventMessage != 0
}

// EventType uses bitmap to convey multiple meanings
type EventType int16

const (
	EventUsage EventType = 1 << iota
	EventMessage
)
