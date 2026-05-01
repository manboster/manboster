package telegram

import (
	"unicode/utf8"

	"github.com/manboster/manboster/internal/util"
)

func (s *Service) Converter(text string, isThinking bool, isToolCall bool) string {
	text, err := util.EscapeMarkdownToTelegramHTML(text)
	if err != nil {
		return text
	}
	if isThinking {
		text = "Model Thinking: \n<blockquote expandable>" + text + "</blockquote>"
	}
	if utf8.RuneCountInString(text) > int(s.cfg.CollapseMsgLength) && !isToolCall {
		text = "<blockquote expandable>" + text + "</blockquote>"
	}
	return text
}
