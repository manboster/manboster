package browser

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/go-rod/rod"
	"github.com/go-rod/stealth"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
)

func (s *Service) BasicScrap(ctx context.Context, url string, respType ResponseType) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err

	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			color.Yellow(i18n.T(keys.BrowserLogCloseBodyFailed))
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	data := s.purgeData(string(body), respType)
	return data, nil
}

func (s *Service) BrowserScrap(ctx context.Context, url string, respType ResponseType, sid string) (string, error) {
	browserInstance, err := s.Manager.getBrowserInstance(ctx, sid)
	if err != nil {
		return "", err
	}

	p, err := stealth.Page(browserInstance.browser)
	if err != nil {
		return "", err
	}
	defer func(p *rod.Page) {
		err := p.Close()
		if err != nil {
			color.Yellow(i18n.T(keys.BrowserLogClosePageFailed))
		}
	}(p)

	err = p.Navigate(url)
	if err != nil {
		return "", err
	}

	err = p.WaitIdle(2 * time.Minute)
	if err != nil {
		return "", err
	}

	str, err := p.HTML()
	if err != nil {
		return "", err
	}

	return s.purgeData(str, respType), nil
}
