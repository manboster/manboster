package gguf

import (
	"context"

	"github.com/hybridgroup/yzma/pkg/llama"
)

func (s *Service) Prepare(ctx context.Context) error {
	libraryPath := libPath()
	mPath, err := modelPath(s.cfg.GGUFurl)
	if err != nil {
		return err
	}

	err = llama.Load(libraryPath)
	if err != nil {
		return err
	}
	llama.LogSet(llama.LogSilent())
	llama.Init()

	model, err := llama.ModelLoadFromFile(mPath, llama.ModelDefaultParams())
	if err != nil {
		return err
	}
	s.model = model

	modelCtx, err := llama.InitFromModel(model, llama.ContextDefaultParams())
	if err != nil {
		return err
	}
	s.modelCtx = modelCtx

	return nil
}
