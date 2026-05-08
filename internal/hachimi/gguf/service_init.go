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
	s.ready = make(chan struct{})

	err = os.MkdirAll(libPath(), 0755)
	if err != nil {
		return err
	}
	err = os.MkdirAll(modelPath(), 0755)
	if err != nil {
		return err
	}

	libInstallPath := libPath()
	if !download.AlreadyInstalled(libInstallPath) {
		s.manager.SetAvail(false)
		go func() {
			err := s.DownloadLibraryRunner(ctx, libInstallPath)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Hachimi Provider] Failed to download library: %s", err))
			}
			s.manager.SetAvail(true)
			s.ready <- struct{}{}
		}()
	} else {
		s.manager.SetAvail(true)
	}

	return s.CheckModel(ctx)
}
