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

func (r ResponseStatusType) String() string {
	switch r {
	case ResponseStatusUnsafe:
		return "Unsafe"
	case ResponseStatusInspect:
		return "Inspect"
	case ResponseStatusSafe:
		return "Safe"
	default:
		return "Unsafe"
	}
}
