package util

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

// EscapeMarkdownToTelegramHTML converts Markdown to Telegram HTML format.
func EscapeMarkdownToTelegramHTML(md string) (string, error) {
	gm := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)

	var buf bytes.Buffer
	if err := gm.Convert([]byte(md), &buf); err != nil {
		return "", err
	}

	htmlStr := buf.String()

	// Convert tables to <pre> blocks
	htmlStr = TableToMarkdown(htmlStr)

	// Parse with goquery for structural processing
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlStr))
	if err != nil {
		return fallbackClean(htmlStr), nil
	}

	// 1. Handle images: Telegram HTML doesn't support <img>
	doc.Find("img").Each(func(_ int, img *goquery.Selection) {
		src, _ := img.Attr("src")
		alt, _ := img.Attr("alt")
		switch {
		case src == "" && alt == "":
			img.ReplaceWithHtml("[image]")
		case src == "":
			img.ReplaceWithHtml(alt)
		case alt == "":
			img.ReplaceWithHtml(src)
		default:
			img.ReplaceWithHtml(fmt.Sprintf("[%s](%s)", alt, src))
		}
	})

	// 2. Clean <a> tags: remove extraneous attributes
	doc.Find("a").Each(func(_ int, a *goquery.Selection) {
		href, exists := a.Attr("href")
		if !exists || href == "" {
			a.ReplaceWithHtml(a.Text())
			return
		}
		inner, _ := a.Html()
		a.ReplaceWithHtml(fmt.Sprintf(`<a href="%s">%s</a>`, href, inner))
	})

	// 3. Fix <pre><code class="language-xxx"> — Telegram HTML doesn't accept class on <code>
	doc.Find("pre code").Each(func(_ int, code *goquery.Selection) {
		// Remove the class attribute entirely
		code.RemoveAttr("class")
	})

	// 4. Fix standalone <code> (inline code) — also strip attributes
	doc.Find("code").Each(func(_ int, code *goquery.Selection) {
		code.RemoveAttr("class")
	})

	// 5. Convert lists — process from innermost to outermost
	for {
		nested := doc.Find("ul, ol")
		if nested.Length() == 0 {
			break
		}
		nested.Each(func(_ int, list *goquery.Selection) {
			if list.Find("ul, ol").Length() > 0 {
				return
			}

			isOrdered := list.Is("ol")
			depth := listDepth(list)
			indent := strings.Repeat("  ", depth)
			var items []string

			list.ChildrenFiltered("li").Each(func(idx int, li *goquery.Selection) {
				text := strings.TrimSpace(li.Text())
				if text == "" {
					return
				}
				if isOrdered {
					items = append(items, indent+fmt.Sprintf("%d. ", idx+1)+text)
				} else {
					items = append(items, indent+"• "+text)
				}
			})

			list.ReplaceWithHtml(strings.Join(items, "\n"))
		})
	}

	// 6. Convert <sub>/<sup>: Telegram doesn't support them, keep text
	doc.Find("sub, sup").Each(func(_ int, el *goquery.Selection) {
		el.ReplaceWithHtml(el.Text())
	})

	// 7. Convert <p> to newline separation
	doc.Find("p").Each(func(_ int, p *goquery.Selection) {
		inner, _ := p.Html()
		p.ReplaceWithHtml(inner)
	})

	// 8. Rename tags to Telegram-compatible versions
	doc.Find("strong").Each(func(_ int, el *goquery.Selection) {
		inner, _ := el.Html()
		el.ReplaceWithHtml("<b>" + inner + "</b>")
	})
	doc.Find("em").Each(func(_ int, el *goquery.Selection) {
		inner, _ := el.Html()
		el.ReplaceWithHtml("<i>" + inner + "</i>")
	})
	doc.Find("del").Each(func(_ int, el *goquery.Selection) {
		inner, _ := el.Html()
		el.ReplaceWithHtml("<s>" + inner + "</s>")
	})
	doc.Find("h1, h2, h3, h4, h5, h6").Each(func(_ int, el *goquery.Selection) {
		inner, _ := el.Html()
		el.ReplaceWithHtml("<b>" + inner + "</b>")
	})

	// 9. Convert <hr> to line separator
	doc.Find("hr").Each(func(_ int, el *goquery.Selection) {
		el.ReplaceWithHtml("\n—————\n")
	})

	// find br and convert it to enter
	doc.Find("br").Each(func(_ int, el *goquery.Selection) {
		el.ReplaceWithHtml("\n\n")
	})

	result, err := doc.Find("body").Html()
	if err != nil {
		return fallbackClean(htmlStr), nil
	}

	// 10. Strip any HTML tags not supported by Telegram
	allowedTags := map[string]bool{
		"b": true, "strong": true,
		"i": true, "em": true,
		"u": true, "ins": true,
		"s": true, "strike": true, "del": true,
		"span":       true, // tg-spoiler
		"tg-spoiler": true,
		"a":          true,
		"tg-emoji":   true,
		"tg-time":    true,
		"code":       true,
		"pre":        true,
		"blockquote": true,
	}
	reStrip := regexp.MustCompile(`</?(\w+)[^>]*>`)
	result = reStrip.ReplaceAllStringFunc(result, func(match string) string {
		reTag := regexp.MustCompile(`</?(\w+)`)
		matches := reTag.FindStringSubmatch(match)
		if len(matches) > 1 && allowedTags[strings.ToLower(matches[1])] {
			return match // keep allowed tags
		}
		// Escape unknown tags so they show as plain text
		return strings.ReplaceAll(strings.ReplaceAll(match, "<", "&lt;"), ">", "&gt;")
	})

	return strings.TrimSpace(result), nil
}

// listDepth calculates nesting depth of a list element
func listDepth(sel *goquery.Selection) int {
	depth := 0
	for parent := sel.Parent(); parent.Length() > 0; parent = parent.Parent() {
		if parent.Is("li") {
			depth++
		}
	}
	return depth
}

// fallbackClean does basic tag replacement if goquery fails
func fallbackClean(htmlStr string) string {
	replacer := strings.NewReplacer(
		"<hr />", "\n—————\n",
		"<hr>", "\n—————\n",
		"<hr/>", "\n—————\n",
		"<h1>", "<b>", "</h1>", "</b>",
		"<h2>", "<b>", "</h2>", "</b>",
		"<h3>", "<b>", "</h3>", "</b>",
		"<h4>", "<b>", "</h4>", "</b>",
		"<h5>", "<b>", "</h5>", "</b>",
		"<h6>", "<b>", "</h6>", "</b>",
		"<sub>", "", "</sub>", "",
		"<sup>", "", "</sup>", "",
		"<p>", "", "</p>", "",
		"<strong>", "<b>", "</strong>", "</b>",
		"<em>", "<i>", "</em>", "</i>",
		"<del>", "<s>", "</del>", "</s>",
		"<br>", "\n", "<br/>", "\n",
	)
	return strings.TrimSpace(replacer.Replace(htmlStr))
}
