package chat

type SelectionCallbackPayload struct {
	SelectionSessionId string // Optional. Required when MessageType = MessageSelectionCallback, it should have a value.
	SelectionValue     string // Optional. Required when MessageType = MessageSelectionCallback, it should be the value of the selection.
}

type SelectionPayload struct {
	Selection   []Selection // Optional. Required when MessageType = MessageSelection, it defines selection's data.
	SelectionId string
}

// Selection provides options to select, answer should be the value.
type Selection struct {
	Name  string // display name
	Value string // actual value
}
