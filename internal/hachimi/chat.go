package hachimi

type Response struct {
	Type   ResponseStatusType
	Reason string
}

type ResponseStatusType int8

const (
	ResponseStatusUnsafe ResponseStatusType = iota
	ResponseStatusInspect
	ResponseStatusSafe
)
