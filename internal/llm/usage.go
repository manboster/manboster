package llm

// Usage is the payload of llm tokens
type Usage struct {
	PromptTokens uint64 // Prompt & input tokens, get from app
	OutputTokens uint64 // The Completion API's output tokens
}
