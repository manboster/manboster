package memory_kv

type NameType string

const (
	NameGet    NameType = "get"
	NameSet    NameType = "set"
	NameList   NameType = "list"
	NameDelete NameType = "delete"
)
