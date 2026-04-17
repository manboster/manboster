package chat

type TextPayload struct {
	Text string // Optional. Required when MessageType = MessageText Text Info
}

type SelectionCallbackPayload struct {
	SelectionSessionId string // Optional. Required when MessageType = MessageSelectionCallback, it should have a value.
	SelectionValue     string // Optional. Required when MessageType = MessageSelectionCallback, it should be the value of the selection.
}

type SelectionPayload struct {
	Selection   []Selection // Optional. Required when MessageType = MessageSelection, it defines selection's data.
	SelectionId string
}

// CommandPayload defines what's in commands
type CommandPayload struct {
	CommandType CommandType // Optional. Required when MessageType = MessageCommand Command's type
	CommandArgs []string    // Optional. Required when MessageType = MessageCommand Command's args
}

// Selection provides options to select, answer should be the value.
type Selection struct {
	Name  string // display name
	Value string // actual value
}
