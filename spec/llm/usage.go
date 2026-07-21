package llm

// Usage is the payload of llm tokens
type Usage struct {
	PromptTokens     int // Prompt & input tokens, get from app
	CompletionTokens int // The Completion API's output tokens
	CachedTokens     int // The token cached in the request
	TotalTokens      int // Total Tokens
	InputCost        float64
	CachedCost       float64
	OutputCost       float64
	TotalCost        float64
}
