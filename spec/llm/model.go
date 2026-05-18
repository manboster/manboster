package llm

// Model defines an information of a LLM model
type Model struct {
	DisplayName     string       `yaml:"display_name" json:"display_name" mapstructure:"display_name"`                // the model's display name
	Name            string       `yaml:"name" json:"name" mapstructure:"name"`                                        // actual name in configuration
	Context         uint64       `yaml:"context" json:"context" mapstructure:"context"`                               // max token of this context
	MaxOutputTokens uint64       `yaml:"max_output_tokens" json:"max_output_tokens" mapstructure:"max_output_tokens"` // max output tokens
	Capabilities    Capabilities `yaml:"capabilities" json:"capabilities" mapstructure:"capabilities"`                // Shows this model's input & output capabilities

	InputPrice  float64 `yaml:"input_price" json:"input_price" mapstructure:"input_price"`    // Optional. Input price USD, per 1m tokens
	OutputPrice float64 `yaml:"output_price" json:"output_price" mapstructure:"output_price"` // Optional. Output price, USD, per 1m tokens
}

// Capabilities defines the capability now using
type Capabilities struct {
	Input  CapabilityType `yaml:"input" json:"input" mapstructure:"input"`
	Output CapabilityType `yaml:"output" json:"output" mapstructure:"output"`
}

// CapabilityType defines the capability this model shows
type CapabilityType int

const (
	CapabilityTextOnly CapabilityType = 1 << iota
	CapabilityToolCall
	CapabilityImage
	CapabilityAudio
	CapabilityVideo
	CapabilityFile
)

const CapabilityTextAndImage = CapabilityText | CapabilityImage
const CapabilityText = CapabilityToolCall | CapabilityTextOnly
const CapabilityAll = CapabilityText | CapabilityImage | CapabilityAudio | CapabilityVideo | CapabilityFile
