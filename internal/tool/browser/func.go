package browser

import (
	"context"
	"errors"

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

func (s *Service) ScrapWebpage(ctx context.Context, url string, effort ScrapType, respType ResponseType) (string, error) {
	switch effort {
	case ScrapTypeText:
		return s.BasicScrap(ctx, url, respType)
	default:
		return "", errors.New("unsupported scrap type")
	}
}

func (s *Service) doWebSearch(ctx context.Context, keyword string, searchEngine EngineType, respType ResponseType) (string, error) {
	return "", nil
}
