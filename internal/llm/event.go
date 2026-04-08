package llm

// Event defines payloads of responses
type Event struct {
	EventType EventType // Required.

	Message *Message // Optional. Required when EventType & EventUsage != 0
	Usage   *Usage   // Optional. Required when EventType & EventMessage != 0
}

// EventType uses bitmap to convey multiple meanings
type EventType int16

const (
	EventUnknown EventType = 1 << iota
	EventUsage
	EventMessage
)
