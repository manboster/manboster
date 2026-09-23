package agent

import "github.com/manboster/manboster/spec/llm"

// CreateQuick create a disposable agent for single use, only one chat time, string returns response message
func CreateQuick() (llm.Message, error) {
	return llm.Message{}, nil
}
