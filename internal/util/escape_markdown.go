package util

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

// EscapeMarkdown helps disable Markdown indicators.
func EscapeMarkdown(text string) string {
	text = strings.ReplaceAll(text, "_", "\\_")
	text = strings.ReplaceAll(text, "*", "\\*")
	text = strings.ReplaceAll(text, "`", "\\`")
	text = strings.ReplaceAll(text, "~", "\\~")
	return text
}

// EscapeMarkdownToTelegramHTML allows Markdown text to send messages in Telegram by using HTML mode. (With help of Gemini)
func EscapeMarkdownToTelegramHTML(md string) (string, error) {
	// configure goldmark
	gm := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithRendererOptions(
			html.WithUnsafe(), // basic HTML is allowed
		),
	)

	var buf bytes.Buffer
	if err := gm.Convert([]byte(md), &buf); err != nil {
		return "", err
	}

	htmlStr := buf.String()

	htmlStr = TableToMarkdown(htmlStr)

	reList := regexp.MustCompile(`(?i)<(ul|ol)[^>]*>`)
	htmlStr = reList.ReplaceAllString(htmlStr, "")

	reListEnd := regexp.MustCompile(`(?i)</(ul|ol)>`)
	htmlStr = reListEnd.ReplaceAllString(htmlStr, "")

	reLi := regexp.MustCompile(`(?i)<li[^>]*>`)
	htmlStr = reLi.ReplaceAllString(htmlStr, "• ")

	reLiEnd := regexp.MustCompile(`(?i)</li>`)
	htmlStr = reLiEnd.ReplaceAllString(htmlStr, "")

	// change tags
	replacer := strings.NewReplacer(
		// hr
		"<hr />", "\n—————\n",
		"<hr>", "\n—————\n",
		"<hr/>", "\n—————\n",
		// titles
		"<h1>", "<b>", "</h1>", "</b>",
		"<h2>", "<b>", "</h2>", "</b>",
		"<h3>", "<b>", "</h3>", "</b>",
		"<h4>", "<b>", "</h4>", "</b>",
		"<h5>", "<b>", "</h5>", "</b>",
		"<h6>", "<b>", "</h6>", "</b>",
		// p
		"<p>", "", "</p>", "\u200B",
		// bold & italic
		"<strong>", "<b>", "</strong>", "</b>",
		"<em>", "<i>", "</em>", "</i>",
		"<del>", "<s>", "</del>", "</s>",
		"<br>", "\n", "<br/>", "\n",
		"<nil>", "[No Response, Go: nil]",
	)

	formatted := replacer.Replace(htmlStr)

	// trim and delete it
	return strings.TrimSpace(formatted), nil
}
