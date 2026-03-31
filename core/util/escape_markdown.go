package util

import "strings"

// EscapeMarkdown helps disable Markdown indicators.
func EscapeMarkdown(text string) string {
	text = strings.ReplaceAll(text, "_", "\\_")
	text = strings.ReplaceAll(text, "*", "\\*")
	text = strings.ReplaceAll(text, "`", "\\`")
	text = strings.ReplaceAll(text, "~", "\\~")
	return text
}
