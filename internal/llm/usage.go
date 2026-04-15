package llm

// Usage is the payload of llm tokens
type Usage struct {
	PromptTokens     int // Prompt & input tokens, get from app
	CompletionTokens int // The Completion API's output tokens
	TotalTokens      int // Total Tokens
}
