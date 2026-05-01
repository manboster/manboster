package browser

import (
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
