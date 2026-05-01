package browser

import (
	"context"
	"io"
	"net/http"

	"github.com/fatih/color"
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
			color.Yellow("[Manboster Tool Provider] We could not close request the body when getting the content")
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

}
