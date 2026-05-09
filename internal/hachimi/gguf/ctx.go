package gguf

type ModelCtx struct {
	NCtx    uint32
	NBatch  uint32
	NUBatch uint32
}

type ModelCtxType string

const (
	ModelCtxLow    ModelCtxType = "low"
	ModelCtxMedium ModelCtxType = "medium"
	ModelCtxHigh   ModelCtxType = "high"
	ModelCtxXHigh  ModelCtxType = "x-high"
)

var ModelCtxMap = map[string]ModelCtx{
	"low": ModelCtx{
		NCtx:    1024,
		NBatch:  768,
		NUBatch: 256,
	},
	"medium": ModelCtx{
		NCtx:    2048,
		NBatch:  1920,
		NUBatch: 256,
	},
	"high": ModelCtx{
		NCtx:    4096,
		NBatch:  3968,
		NUBatch: 256,
	},
	"x-high": ModelCtx{
		NCtx:    8192,
		NBatch:  7936,
		NUBatch: 256,
	},
}
