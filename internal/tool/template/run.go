package template

import (
	"context"

	"github.com/manboster/manboster/spec/plugin"
)

func (s *Service) Init(ctx context.Context, cfg any) error {
	return nil
}

func (s *Service) Start(ctx context.Context) error {
	return nil
}

func (s *Service) Run(ctx context.Context, args string) (*plugin.RunResponse, error) {
	return nil, nil
}

func (s *Service) Continue(ctx context.Context, session string) (*plugin.RunResponse, error) {
	return nil, nil
}

func (s *Service) Stop() error {
	return nil
}
