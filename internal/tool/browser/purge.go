package browser

import (
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
)

func (s *Service) purgeData(respStr string, respType ResponseType) string {
	switch respType {
	case ResponseTypeRaw:
		return respStr
	case ResponseTypeBody:
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(respStr))
		if err != nil {
			color.Yellow(i18n.T(keys.BrowserLogPurgeOpenDocFailed))
			return respStr
		}
		html, err := doc.Find("body").Html()
		if err != nil {
			color.Yellow(i18n.T(keys.BrowserLogPurgeBodyFailed))
			return respStr
		}
		return html
	case ResponseTypeMarkdown:
		conv := md.NewConverter("", true, nil)
		markdown, err := conv.ConvertString(respStr)
		if err != nil {
			color.Yellow(i18n.T(keys.BrowserLogPurgeMarkdownFailed))
			return respStr
		}
		return markdown
	default:
		return respStr
	}
}
