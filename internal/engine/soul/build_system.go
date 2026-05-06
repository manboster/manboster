package soul

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config/prompt"
	"github.com/manboster/manboster/spec/llm"
)

var re = regexp.MustCompile(`(?s)## Tone and Formatting.*?## Reminders`)

// BuildSystemMessage returns system prompt message
func (s *Service) BuildSystemMessage(ctx context.Context, souls []string) (llm.Message, error) {
	var text strings.Builder
	for _, soul := range souls {
		// system soul is automatically global!
		if soul == "system" {
			basePrompt := prompt.InitialSystemPrompt
			if so, avail := s.soulMap["system"]; avail {
				replacement := fmt.Sprintf("# Tone and Formatting\n%s\n", so.Content)
				basePrompt = re.ReplaceAllString(basePrompt, replacement)
			}
			text.WriteString(basePrompt + "\n")
		} else {
			so, avail := s.soulMap[soul]
			if !avail {
				color.Yellow(fmt.Sprintf("[Manboster Soul] soul %s not available", soul))
				continue
			}
			text.WriteString(so.Content + "\n")
		}
	}
	return llm.Message{
		Type: llm.MessageText,
		Role: llm.RoleSystem,
		Parts: []llm.MessageParts{
			{
				PartsType: llm.MessagePartsText,
				Text: &llm.MessageTextPayload{
					Text: text.String(),
				},
			},
		},
	}, nil
}
