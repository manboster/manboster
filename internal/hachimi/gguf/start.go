package gguf

import (
	"context"

	"github.com/fatih/color"
	"github.com/hybridgroup/yzma/pkg/llama"
)

func (s *Service) Start(ctx context.Context) error {
	if !s.manager.IsReady() {
		return s.CheckReadyRunner(ctx)
	}

	go func() {
		err := s.GCRunner(ctx)
		if err != nil {
			color.Yellow("[Manboster Hachimi Provider] Failed to start gc runner!")
		}
	}()

	return s.Prepare(ctx)
}

func (s *Service) Stop() error {
	defer func() {
		if r := recover(); r != nil {
			// close may be unstable so there is nothing available!
			// color.Yellow(fmt.Sprintf("[Manboster Downloader] Recovered in %v", r))
			return
		}
	}()

	if s.ready != nil {
		close(s.ready)
	}
	if llama.Close != nil {
		llama.Close()
	}
	if llama.Free != nil {
		err := llama.Free(s.manager.ModelCtx())
		if err != nil {
			return err
		}
	}
	return nil
}
