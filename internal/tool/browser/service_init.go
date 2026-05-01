package browser

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/go-rod/rod/lib/launcher"
)

func (s *Service) Init(ctx context.Context, cfg any) error {
	browserManager := launcher.NewBrowser()
	_, err := browserManager.Get()
	if err != nil {
		go func() {
			err := s.DownloadBrowser(browserManager)
			if err != nil {
				color.Red(fmt.Sprintf("[Manboster Tool Provider] Failed to download browser: %v", err))
			}
		}()
	} else {
		s.isReady = true
	}
	return nil
}
