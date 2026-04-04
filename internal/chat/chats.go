package chat

// ChatsType is an enum defining chat types.
type ChatsType int16

const (
	ChatsUnknown  ChatsType = 0
	ChatsPersonal ChatsType = 1
	ChatsGroup    ChatsType = 2
	ChatsChannel  ChatsType = 3
)
