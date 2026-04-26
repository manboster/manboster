package llm

type MessageParts struct {
	PartsType MessagePartsType
	Text      *MessageTextPayload
}

type MessagePartsType int8

const (
	MessagePartsText MessagePartsType = iota
	MessagePartsImage
	MessagePartsFile
)
