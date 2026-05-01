package browser

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/go-rod/rod/lib/launcher"
)

func (s *Service) Init(ctx context.Context, cfg any) error {
	go func() {
		browserManager := launcher.NewBrowser()
		_, err := browserManager.Get()
		if err != nil {
			err := s.DownloadBrowser(browserManager)
			if err != nil {
				color.Red(fmt.Sprintf("[Manboster Tool Provider] Failed to download browser: %v", err))
			}
		} else {
			s.isReady = true
		}
	}()
	s.Manager = NewManager()
	return nil
}
