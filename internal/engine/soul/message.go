package soul

import (
	"strings"

	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/llm"
)

// ChatMessageToString converts a chat's message to string
func (s *Service) ChatMessageToString(msg *chat.Message) string {
	var respString strings.Builder
	if msg.MessageType&chat.MessageText != 0 && msg.Text != nil {
		respString.WriteString(msg.Text.Text)
	}

	// TODO: images, files, and more...
	return respString.String()
}

// LLMMessageToString TODO: converts a LLM message to string format
func (s *Service) LLMMessageToString(msg *llm.Message) (*chat.Message, string) {
	return &chat.Message{}, ""
}
