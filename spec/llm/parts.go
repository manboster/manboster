package llm

type MessageParts struct {
	PartsType MessagePartsType
	Text      *MessageTextPayload
	Image     *MessageImagePayload
	File      *MessageFilePayload
}

type MessagePartsType int8

const (
	MessagePartsText MessagePartsType = iota
	MessagePartsImage
	MessagePartsFile
)
