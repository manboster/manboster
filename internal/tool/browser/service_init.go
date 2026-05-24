package browser

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-viper/mapstructure/v2"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
)

func (s *Service) Init(ctx context.Context, conf any) error {
	if svc != nil {
		return nil
	}

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
				color.Red(fmt.Sprintf(i18n.T(keys.BrowserLogDownloadFailed), err))
			}
		} else {
			s.isReady = true
		}
	}()

	s.Manager = NewManager(s.cfg)
	svc = s
	return nil
}
