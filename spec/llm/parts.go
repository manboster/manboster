package llm

type MessageParts struct {
	PartsType MessagePartsType
	Text      *MessageTextPayload
	Image     *MessageImagePayload
}

type MessagePartsType int8

const (
	MessagePartsText MessagePartsType = iota
	MessagePartsImage
	MessagePartsFile
)
