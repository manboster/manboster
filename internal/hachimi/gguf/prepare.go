package gguf

import (
	"context"

	"github.com/fatih/color"
	"github.com/hybridgroup/yzma/pkg/llama"
)

func (s *Service) Prepare(ctx context.Context) error {
	color.Blue("[Manboster Hachimi Provider] Preparing Hachimi Model...")
	libraryPath := libPath()

	err := llama.Load(libraryPath)
	if err != nil {
		s.manager.SetAvail(false)
		return err
	}
	llama.LogSet(llama.LogSilent())
	llama.Init()
	return nil
}

func (s *Service) LoadModel(ctx context.Context) error {
	mPath, err := modelFilePath(s.cfg.GGUFurl)
	if err != nil {
		return err
	}
	model, err := llama.ModelLoadFromFile(mPath, llama.ModelDefaultParams())
	if err != nil {
		return err
	}
	s.manager.SetModel(model)

	ctxParams := llama.ContextDefaultParams()
	ctxParams.NCtx = 2048
	ctxParams.NBatch = 512
	ctxParams.NUbatch = 256

	modelCtx, err := llama.InitFromModel(model, ctxParams)
	if err != nil {
		return err
	}
	s.manager.SetModelCtx(modelCtx)
	color.Green("[Manboster Hachimi Provider] Hachimi Model loaded from memory!")

	template := llama.ModelChatTemplate(model, "")
	if template == "" {
		template = "chatml"
	}
	s.chatTemplate = template

	sampler := llama.SamplerChainInit(llama.SamplerChainDefaultParams())
	llama.SamplerChainAdd(sampler, llama.SamplerInitGreedy())
	s.sampler = sampler

	vocab := llama.ModelGetVocab(s.manager.Model())
	s.vocab = vocab

	color.Blue("[Manboster Hachimi Provider] Ready to go!")

	s.manager.SetLoad(true)
	return nil
}

func (s *Service) FreeModel() error {
	err := llama.Free(s.manager.ModelCtx())
	if err != nil {
		return err
	}
	err = llama.ModelFree(s.manager.Model())
	if err != nil {
		return err
	}
	s.manager.SetLoad(false)
	return nil
}
