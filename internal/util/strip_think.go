package util

import "regexp"

// StripThink strips think content to another message.
func StripThink(input string) string {
	re := regexp.MustCompile(`(?s)<think>.*?</think>`)
	return re.ReplaceAllString(input, "")
}

// ExtractThinkContent extracts thinking content.
func ExtractThinkContent(input string) string {
	re := regexp.MustCompile(`(?s)<think>(.*?)</think>`)
	match := re.FindStringSubmatch(input)

	if len(match) > 1 {
		return match[1] // raw content
	}
	return ""
}
