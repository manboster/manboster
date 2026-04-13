package chat

type AbilityType int16

const (
	AbilitySendText AbilityType = 1 << iota
	AbilitySendImage
	AbilitySendVoice
	AbilitySendVideo
	AbilitySendFile
	AbilitySendSelect
)
