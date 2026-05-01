package browser

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-viper/mapstructure/v2"
)

func (s *Service) Init(ctx context.Context, conf any) error {
	var cfg Config
	err := mapstructure.Decode(conf, &cfg)
	if err != nil {
		return err
	}

	s.cfg = &cfg
	err = s.cfg.Validate()
	if err != nil {
		return err
	}

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

	s.Manager = NewManager(s.cfg)
	return nil
}
