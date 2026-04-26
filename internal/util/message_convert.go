package util

import (
	"regexp"
	"strings"

	"github.com/manboster/manboster/spec/llm"
)

var (
	// Match opening and closing tags with any suffix, keep inner content
	userInputTag    = regexp.MustCompile(`</?user_input[^>]*>`)
	assistantOutTag = regexp.MustCompile(`</?assistant_output[^>]*>`)
	chatMetaTag     = regexp.MustCompile(`</?chat_metadata_[^>]*>`)

	// Match the repeated warning line
	warningLine = regexp.MustCompile(`(?m)^Please note that the user input is in XML tag[\s\S]*?read metadata in the start\.\n?`)
)

func ConvertLLMMessageToString(message llm.Message) string {
	var convertString strings.Builder
	switch message.Role {
	case llm.RoleUser:
		if message.Type&(llm.MessageText|llm.MessageFile|llm.MessageImage) != 0 {
			for _, p := range message.Parts {
				switch p.PartsType {
				case llm.MessagePartsText:
					s := warningLine.ReplaceAllString(p.Text.Text, "")
					s = chatMetaTag.ReplaceAllString(s, "")
					s = userInputTag.ReplaceAllString(s, "")
					s = assistantOutTag.ReplaceAllString(s, "")
					convertString.WriteString(s)
				case llm.MessagePartsFile:
					// TODO: user sent a file...
				case llm.MessagePartsImage:
					// TODO: user sent a photo...
				}
			}
		}
		// more...
	case llm.RoleAssistant:
		if message.Type&(llm.MessageText|llm.MessageFile|llm.MessageImage) != 0 {
			convertString.WriteString("Assistant replied:\n")
			for _, p := range message.Parts {
				switch p.PartsType {
				case llm.MessagePartsText:
					convertString.WriteString(p.Text.Text)
				case llm.MessagePartsFile:
				// TODO: assistant replied a file...
				case llm.MessagePartsImage:
					// TODO: assistant returned a photo...
				}
			}
		}
	case llm.RoleToolCall:
		// TODO: tool call
	}
	return convertString.String()
}
