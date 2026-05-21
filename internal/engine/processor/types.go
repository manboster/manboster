package processor

type ProcessSuggestion int8

const (
	ProcessHandle   ProcessSuggestion = iota
	ProcessDrop     ProcessSuggestion = iota
	ProcessConsider ProcessSuggestion = iota
)

func (p ProcessSuggestion) String() string {
	switch p {
	case ProcessHandle:
		return "Handle"
	case ProcessDrop:
		return "Drop"
	case ProcessConsider:
		return "Consider"
	default:
		return "Unknown"
	}
}
