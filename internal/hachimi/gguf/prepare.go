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

	ctxParams := llama.ContextDefaultParams()
	ctxParams.NCtx = 8192
	ctxParams.NBatch = 1024
	ctxParams.NUbatch = 1024

	modelCtx, err := llama.InitFromModel(model, ctxParams)
	if err != nil {
		return err
	}
	s.modelCtx = modelCtx

	sampler := llama.SamplerChainInit(llama.SamplerChainDefaultParams())
	llama.SamplerChainAdd(sampler, llama.SamplerInitGreedy())
	s.sampler = sampler

	vocab := llama.ModelGetVocab(s.model)
	s.vocab = vocab

	s.ready <- struct{}{}
	return nil
}
