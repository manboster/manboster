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

	s.manager = NewManager()

	libInstallPath := libPath()
	if !download.AlreadyInstalled(libInstallPath) {
		s.manager.SetAvail(false)
		go func() {
			err := s.DownloadLibraryRunner(libInstallPath)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Hachimi Provider] Failed to download library: %s", err))
			}
			s.manager.SetAvail(true)
		}()
	} else {
		s.manager.SetAvail(true)
	}

	modelFilePath, err := modelPath(s.cfg.GGUFurl)
	if err != nil {
		return err
	}
	if _, err = os.Stat(modelFilePath); os.IsNotExist(err) {
		s.manager.SetAvailModel(false)
		go func() {
			err := Download(ctx, cfg.GGUFurl, modelFilePath)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Hachimi Provider] Failed to download model: %s", err))
			}
			s.manager.SetAvailModel(true)
		}()
	} else if os.IsExist(err) {
		s.manager.SetAvailModel(true)
	} else {
		return err
	}

	return nil
}
