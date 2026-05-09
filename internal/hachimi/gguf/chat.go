package gguf

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/hybridgroup/yzma/pkg/llama"
	"github.com/manboster/manboster/internal/hachimi"
)

func (s *Service) Chat(ctx context.Context, evalMsg string) (*hachimi.Response, error) {
	s.chatLock.Lock()
	defer s.chatLock.Unlock()

	if !s.manager.IsReady() {
		if !s.manager.Avail() {
			return nil, ErrNotAvailable
		}
	}

	if !s.manager.Load() {
		color.Blue(fmt.Sprintf("[Manboster Hachimi Provider] Model not loaded, loading..."))
		err := s.LoadModel(ctx)
		if err != nil {
			return nil, err
		}
	}

	s.lastUse = time.Now()

	messages := make([]llama.ChatMessage, 0)
	switch s.cfg.ModelType {
	case ModelLLM:
		messages = append(messages, llama.NewChatMessage("system", safetyLLMSystemPrompt))
	case ModelSafeguard:
		messages = append(messages, llama.NewChatMessage("system", safetySafeguardSystemPrompt))
	}
	messages = append(messages, llama.NewChatMessage("user", evalMsg))

	buf := make([]byte, 4096)
	n := llama.ChatApplyTemplate(s.chatTemplate, messages, true, buf)
	if n <= 0 {
		s.manager.SetAvail(false)
		return nil, fmt.Errorf("failed to apply chat template")
	}
	if n > 3968 {
		return nil, nil
	}
	prompt := string(buf[:n])

	// color.Blue(prompt)
	tokens := llama.Tokenize(s.vocab, prompt, true, true)
	batch := llama.BatchGetOne(tokens)

	resp := ""

	if llama.ModelHasEncoder(s.manager.Model()) {
		_, err := llama.Encode(s.manager.ModelCtx(), batch)
		if err != nil {
			s.manager.SetAvail(false)
			return nil, err
		}

		start := llama.ModelDecoderStartToken(s.manager.Model())
		if start == llama.TokenNull {
			start = llama.VocabBOS(s.vocab)
		}

		batch = llama.BatchGetOne([]llama.Token{start})

	}

	for pos := int32(0); pos < 128; pos += batch.NTokens {
		_, err := llama.Decode(s.manager.ModelCtx(), batch)
		if err != nil {
			s.manager.SetAvail(false)
			return nil, err
		}
		token := llama.SamplerSample(s.sampler, s.manager.ModelCtx(), -1)

		if llama.VocabIsEOG(s.vocab, token) {
			break
		}

		buf := make([]byte, 256)
		l := llama.TokenToPiece(s.vocab, token, buf, 0, false)
		next := string(buf[:l])

		batch = llama.BatchGetOne([]llama.Token{token})

		resp += next
	}

	// color.Blue(resp)

	switch s.cfg.ModelType {
	case ModelLLM:
		return s.purgeLLMChatData(resp)
	case ModelSafeguard:
		return s.purgeSafeguardChatData(resp)
	}
	return nil, nil
}
