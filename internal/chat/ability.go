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

var AbilityAll = AbilityNoSelect | AbilitySendSelect
var AbilityTextAndImage = AbilitySendText | AbilitySendImage
var AbilityTextAndFile = AbilitySendText | AbilitySendFile
var AbilityNoSelect = AbilitySendText | AbilitySendImage | AbilitySendVoice | AbilitySendFile | AbilitySendVideo
