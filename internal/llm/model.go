package llm

import "math"

// Model defines an information of a LLM model
type Model struct {
	DisplayName     string       // the model's display name
	Name            string       // actual name in configuration
	Context         uint64       // max token of this context
	MaxOutputTokens uint64       // max output tokens
	Capabilities    Capabilities // Shows this model's input & output capabilities

	InputPrice  float64 // Optional. Input price USD, per 1m tokens
	OutputPrice float64 // Optional. Output price, USD, per 1m tokens
}

type Capabilities struct {
	Input  CapabilityType
	Output CapabilityType
}

// CapabilityType defines the capability this model shows
type CapabilityType int

const (
	CapabilityText CapabilityType = 1 << iota
	CapabilityImage
	CapabilityAudio
	CapabilityVideo
	CapabilityFile
)

const CapabilityAll = CapabilityText | CapabilityImage | CapabilityAudio | CapabilityVideo
const CapabilityTextAndImage = CapabilityText | CapabilityImage

// CalculateCompactTokens returns when the tokens above which, it should be compacted and open a new conversation
func CalculateCompactTokens(m Model) uint64 {
	if m.Context == 0 {
		return 0
	}
	if m.MaxOutputTokens == 0 || (float64(m.Context)-float64(m.MaxOutputTokens)) < 0 {
		return uint64(math.Floor(float64(m.Context) * 0.6))
	}
	return uint64(math.Floor((float64(m.Context) - float64(m.MaxOutputTokens)) * 0.8))
}
