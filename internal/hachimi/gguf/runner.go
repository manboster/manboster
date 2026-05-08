package gguf

import (
	"context"
	"fmt"
	"runtime"

	"github.com/fatih/color"
	"github.com/hybridgroup/yzma/pkg/download"
)

func (s *Service) DownloadLibraryRunner(path string) error {
	version, err := download.LlamaLatestVersion()
	if err != nil {
		color.Yellow(fmt.Sprintf("[Manboster Hachimi Provider] could not get the latest version of llama.cpp: %q", err))
		return err
	}
	// normally it's cpu
	processor := "cpu"
	// detect there is a cuda environment or not
	if cudaInstalled, cudaVersion := download.HasCUDA(); cudaInstalled {
		color.Yellow(fmt.Sprintf("[Manboster Hachimi Provider] CUDA detected (version %s), using CUDA build.", cudaVersion))
		processor = "cuda"
	}
	// TODO: add mlx

	color.Blue("[Manboster Hachimi Provider] downloading necessary llama.cpp build offline...")
	if err := download.Get(runtime.GOARCH, runtime.GOOS, processor, version, path); err != nil {
		color.Red(fmt.Sprintf("[Manboster Hachimi Provider] Failed to download llama.cpp: %q", err))
		return err
	}
	color.Green("[Manboster Hachimi Provider] successfully downloaded llama.cpp!")
	return nil
}

func (s *Service) CheckReadyRunner(ctx context.Context) error {
	for {
		select {
		case <-s.ready:
			return s.Prepare(ctx)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
