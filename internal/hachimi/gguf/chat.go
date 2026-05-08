package gguf

import (
	"context"

	"github.com/fatih/color"
	"github.com/hybridgroup/yzma/pkg/llama"
	"github.com/manboster/manboster/internal/hachimi"
)

func (s *Service) Chat(ctx context.Context, sysMsg string, evalMsg string) (*hachimi.Response, error) {
	if !s.manager.IsReady() {
		return nil, ErrNotAvailable
	}

	messages := make([]llama.ChatMessage, 0)
	if sysMsg != "" {
		messages = append(messages, llama.NewChatMessage("system", sysMsg))
	}
	messages = append(messages, llama.NewChatMessage("user", evalMsg))

	tokens := llama.Tokenize(s.vocab, evalMsg, true, true)
	batch := llama.BatchGetOne(tokens)

	resp := ""

	if llama.ModelHasDecoder(s.model) {
		_, err := llama.Encode(s.modelCtx, batch)
		if err != nil {
			return nil, err
		}

		start := llama.ModelDecoderStartToken(s.model)
		if start == llama.TokenNull {
			start = llama.VocabBOS(s.vocab)
		}

		batch = llama.BatchGetOne([]llama.Token{start})

	}

	for pos := int32(0); pos < 128; pos += batch.NTokens {
		decode, err := llama.Decode(s.modelCtx, batch)
		if err != nil {
			return nil, err
		}
		token := llama.SamplerSample(s.sampler, s.modelCtx, decode)

		if llama.VocabIsEOG(s.vocab, token) {
			break
		}

		buf := make([]byte, 256)
		l := llama.TokenToPiece(s.vocab, token, buf, 0, false)
		next := string(buf[l:])

		batch = llama.BatchGetOne([]llama.Token{token})

		resp += next
	}

	color.Blue(resp)

	return nil, nil
}
