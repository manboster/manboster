package gguf

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/hybridgroup/yzma/pkg/llama"
)

func (s *Service) Start(ctx context.Context) error {
	if !s.manager.IsReady() {
		go func() {
			err := s.CheckReadyRunner(ctx)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Downloader] Could not check runner: %q", err))
			}
		}()
		return nil
	}

	return s.Prepare(ctx)
}

func (s *Service) Stop() error {
	if s.ready != nil {
		close(s.ready)
	}
	llama.Close()
	err := llama.Free(s.modelCtx)
	if err != nil {
		return err
	}
	return nil
}
