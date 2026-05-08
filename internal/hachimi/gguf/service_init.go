package gguf

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/hybridgroup/yzma/pkg/download"
	"github.com/manboster/manboster/internal/config"
)

func (s *Service) Init(ctx context.Context, conf any) error {
	libInstallPath := config.Path(filepath.Join("hachimi", "llama.cpp"))
	if !download.AlreadyInstalled(libInstallPath) {
		s.avail = false
		go func() {
			err := s.DownloadLibraryRunner()
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Hachimi Provider] Failed to download library: %s", err))
			}
		}()
	} else {
		s.avail = true
	}
	return nil
}
