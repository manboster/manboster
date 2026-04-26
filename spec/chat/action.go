package chat

// ActionType gives you the type of current action's callback.
type ActionType string

const (
	ActionUnknown ActionType = ""
	ActionPending ActionType = "pending" // received request
	ActionSuccess ActionType = "success"
	ActionError   ActionType = "error"
)
