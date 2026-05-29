package hachimi

type Response struct {
	Type   ResponseStatusType
	Reason string
}

type ResponseStatusType int8

const (
	ResponseStatusUnsafe  ResponseStatusType = -1
	ResponseStatusInspect ResponseStatusType = 1
	ResponseStatusSafe    ResponseStatusType = 2
)
