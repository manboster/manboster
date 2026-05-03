package browser

import (
	"context"
	"errors"
	"net/url"

	"github.com/fatih/color"
	"github.com/go-rod/rod/lib/launcher"
)

func (s *Service) DownloadBrowser(m *launcher.Browser) error {
	color.Yellow("[Manboster Tool Provider] Could not get an available browser from your machine, now downloading...")
	err := m.Download()
	if err != nil {
		return err
	}
	s.isReady = true
	return nil
}

func (s *Service) ScrapWebpage(ctx context.Context, url string, effort ScrapType, respType ResponseType, sid string) (string, error) {
	switch effort {
	case ScrapTypeText:
		return s.BasicScrap(ctx, url, respType)
	case ScrapTypeBrowser:
		return s.BrowserScrap(ctx, url, respType, sid)
	default:
		return "", errors.New("unsupported scrap type")
	}
}

func (s *Service) doWebSearch(ctx context.Context, keyword string, searchEngine EngineType, respType ResponseType, sid string) (string, error) {
	u := ""
	switch searchEngine {
	case EngineGoogle:
		u = "https://www.google.com/search?q=" + url.QueryEscape(keyword)
	case EngineBaidu:
		u = "https://www.baidu.com/s?ie=utf-8&wd=" + url.QueryEscape(keyword)
	case EngineBing:
		u = "https://www.bing.com/search?q=" + url.QueryEscape(keyword)
	case EngineCNBing:
		u = "https://cn.bing.com/search?q=" + url.QueryEscape(keyword)
	case EngineDuckDuckGo:
		u = "https://duckduckgo.com/search?q=" + url.QueryEscape(keyword)
	case EngineGitHub:
		u = "https://github.com/search?q=" + url.QueryEscape(keyword)
	case EngineWikipedia:
		u = "https://wikipedia.org/w/index.php?title=Special:Search&fulltext=1&ns0=1&search=" + url.QueryEscape(keyword)
	default:
		return "", errors.New("unsupported engine type")
	}
	str, err := s.BasicScrap(ctx, u, respType)
	if err != nil {
		color.Yellow("[Manboster Tool Provider] We could not get content directly, opening the browser...")
		return "", err
	}
	return str, nil
}
