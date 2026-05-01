package browser

import (
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

func (s *Service) purgeData(respStr string, respType ResponseType) string {
	switch respType {
	case ResponseTypeRaw:
		return respStr
	case ResponseTypeBody:
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(respStr))
		if err != nil {
			color.Yellow("[Manboster Tool Provider] Could not open document when purging data")
			return respStr
		}
		html, err := doc.Find("body").Html()
		if err != nil {
			color.Yellow("[Manboster Tool Provider] Could not find body when purging data")
			return respStr
		}
		return html
	case ResponseTypeMarkdown:
		conv := md.NewConverter("", true, nil)
		markdown, err := conv.ConvertString(respStr)
		if err != nil {
			color.Yellow("[Manboster Tool Provider] Could not convert from html to markdown")
			return respStr
		}
		return markdown
	default:
		return respStr
	}
}
