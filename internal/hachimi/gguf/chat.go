package gguf

import (
	"context"

	"github.com/manboster/manboster/internal/hachimi"
)

func (s *Service) Chat(ctx context.Context, sysMsg string, evalMsg string) (*hachimi.Response, error) {
	if s.avail && s.availModel {
		return nil, ErrNotAvailable
	}
	return nil, nil
}
