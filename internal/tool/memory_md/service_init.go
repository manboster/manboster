package memory_md

import (
	"context"
	"fmt"
	"os"

	"github.com/manboster/manboster/internal/config"
)

func (s *Service) Init(ctx context.Context, cfg any) error {
	if err := os.MkdirAll(config.Path("memory"), 0755); err != nil {
		return fmt.Errorf("failed to create memory dir: %w", err)
	}
	return nil
}

func (s *Service) Start(ctx context.Context) error {
	return nil
}
