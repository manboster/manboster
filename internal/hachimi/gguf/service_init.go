package gguf

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/go-viper/mapstructure/v2"
	"github.com/hybridgroup/yzma/pkg/download"
)

func (s *Service) Init(ctx context.Context, conf any) error {
	var cfg Config
	// get config
	err := mapstructure.Decode(conf, &cfg)
	if err != nil {
		return err
	}

	s.cfg = &cfg
	err = s.cfg.Validate()
	if err != nil {
		return err
	}

	libInstallPath := libPath()
	if !download.AlreadyInstalled(libInstallPath) {
		s.avail = false
		go func() {
			err := s.DownloadLibraryRunner(libInstallPath)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Hachimi Provider] Failed to download library: %s", err))
			}
			s.avail = true
		}()
	} else {
		s.avail = true
	}

	if _, err = os.Stat(modelFilePath); os.IsNotExist(err) {
		s.availModel = false
		go func() {
			err := Download(ctx, cfg.GGUFurl, modelFilePath)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Hachimi Provider] Failed to download model: %s", err))
			}
			s.availModel = true
		}()
	} else if os.IsExist(err) {
		s.availModel = true
	} else {
		return err
	}

	return nil
}
