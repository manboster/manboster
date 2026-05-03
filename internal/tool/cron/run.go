package cron

import (
	"context"

	"github.com/manboster/manboster/spec/plugin"
)

func (s *Service) Run(ctx context.Context, args string) (*plugin.RunResponse, error) {
	return nil, nil
}

func (s *Service) Continue(ctx context.Context, session string) (*plugin.RunResponse, error) {
	return nil, nil
}

func (s *Service) Close() error {
	return nil
}
