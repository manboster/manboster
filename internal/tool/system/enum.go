package system

type NameType string

const (
	NameOSInfo  NameType = "os_info"
	NameProcess NameType = "process"
)

type ActionType string

const (
	ActionList ActionType = "list"
	ActionInfo ActionType = "info"
	ActionKill ActionType = "kill"
)
