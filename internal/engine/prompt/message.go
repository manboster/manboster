package prompt

import (
	"strings"

	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/llm"
)

// ChatMessageToString converts a chat's message to string
func (s *Service) ChatMessageToString(msg *chat.Message) (*llm.Message, string) {
	var m llm.Message
	var respString strings.Builder

	if msg.MessageType&chat.MessageText != 0 && msg.Text != nil {
		m.Type |= llm.MessageText
		m.Parts = append(m.Parts, llm.MessageParts{
			PartsType: llm.MessagePartsText,
			Text: &llm.MessageTextPayload{
				Text: msg.Text.Text,
			},
		})
		respString.WriteString(msg.Text.Text)
	}

	// TODO: images, files, and more...
	return &m, respString.String()
}

// LLMMessageToString TODO: converts a LLM message to string format
func (s *Service) LLMMessageToString(msg *llm.Message) (*chat.Message, string) {
	return &chat.Message{}, ""
}
