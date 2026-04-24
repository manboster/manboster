package util

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// TableToMarkdown converts from HTML to Markdown
func TableToMarkdown(h string) string {
	if !strings.Contains(strings.ToLower(h), "<table") {
		return h
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(h))
	if err != nil {
		return h // failed, return
	}

	// iterate table information
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		var buf bytes.Buffer

		buf.WriteString("<pre>\n")
		table.Find("tr").Each(func(rowIndex int, row *goquery.Selection) {
			var rowData []string
			row.Find("th, td").Each(func(colIndex int, cell *goquery.Selection) {
				text := strings.TrimSpace(cell.Text())
				text = strings.ReplaceAll(text, "\n", " ")
				rowData = append(rowData, text)
			})
			if len(rowData) == 0 {
				return
			}
			buf.WriteString("| " + strings.Join(rowData, " | ") + " |\n")
			if rowIndex == 0 {
				var separator []string
				for range rowData {
					separator = append(separator, "---")
				}
				buf.WriteString("| " + strings.Join(separator, " | ") + " |\n")
			}
		})

		buf.WriteString("</pre>\n")
		table.ReplaceWithHtml(buf.String())
	})

	htmlStr, err := doc.Find("body").Html()
	if err != nil {
		return h
	}

	return strings.TrimSpace(htmlStr)
}
