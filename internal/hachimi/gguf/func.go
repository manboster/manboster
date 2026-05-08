package gguf

import (
	"fmt"
	"net/url"
	"path"
	"path/filepath"

	"github.com/manboster/manboster/internal/config"
)

func libPath() string {
	return config.Path(filepath.Join("hachimi", "llama.cpp"))
}

func modelPath(url string) (string, error) {
	name, err := modelName(url)
	if err != nil {
		return "", err
	}
	modelFilePath := config.Path(filepath.Join("hachimi", "models", name))
	return modelFilePath, nil
}

func modelName(ur string) (string, error) {
	u, err := url.Parse(ur)
	if err != nil {
		return "", err
	}
	modelFileName := path.Base(u.Path)
	if modelFileName == "" || u.Scheme != "http" || u.Scheme != "https" || u.Scheme != "file" {
		return "", fmt.Errorf("invalid URL: %s", u.String())
	}
	return modelFileName, nil
}
