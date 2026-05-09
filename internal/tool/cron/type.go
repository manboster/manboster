package cron

type MessageType string

const (
	MessageText   MessageType = "text"
	MessagePrompt MessageType = "prompt"
)

type ToChatType string

const (
	ToThisChat ToChatType = "this"
	ToPM       ToChatType = "pm"
)

type NameType string

const (
	NameGet    NameType = "get"
	NameSet    NameType = "set"
	NameList   NameType = "list"
	NameDelete NameType = "delete"
)
