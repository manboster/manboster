package gguf

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
)

func libPath() string {
	return config.Path(filepath.Join("hachimi", "llama.cpp"))
}

func modelFilePath(url string) (string, error) {
	name, err := modelName(url)
	if err != nil {
		return "", err
	}
	modelFPath := filepath.Join(modelPath(), name)
	return modelFPath, nil
}

func modelPath() string {
	return config.Path(filepath.Join("hachimi", "models"))
}

func modelName(ur string) (string, error) {
	u, err := url.Parse(ur)
	if err != nil {
		return "", err
	}
	modelFileName := path.Base(u.Path)
	if modelFileName == "" || (u.Scheme != "http" && u.Scheme != "https" && u.Scheme != "file") {
		return "", fmt.Errorf("invalid URL: %s", u.String())
	}
	return modelFileName, nil
}

func (s *Service) CheckModel(ctx context.Context) error {
	modelFileName, err := modelFilePath(s.cfg.GGUFurl)
	if err != nil {
		return err
	}
	color.Blue("[Manboster Hachimi Provider] Loding models...")

	_, err = os.Stat(modelFileName)
	if os.IsNotExist(err) {
		s.manager.SetAvailModel(false)
		go func() {
			color.Yellow("[Manboster Hachimi Provider] Model not found, starting to download...")
			dlErr := Download(ctx, s.cfg.GGUFurl, modelFileName)
			if dlErr != nil {
				color.Yellow(fmt.Sprintf("[Manboster Hachimi Provider] Failed to download model: %s", err))
			}
			s.manager.SetAvailModel(true)
			s.ready <- struct{}{}
		}()
	} else if os.IsExist(err) || err == nil {
		color.Green("[Manboster Hachimi Provider] Model loaded!")
		s.manager.SetAvailModel(true)
	} else {
		return err
	}
	return nil
}
