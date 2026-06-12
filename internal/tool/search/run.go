package search

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/manboster/manboster/spec/plugin"
)

func (s *Service) Init(ctx context.Context, cfg any) error {
	return nil
}

func (s *Service) Start(ctx context.Context) error {
	return nil
}

func (s *Service) Run(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg := RunArgs{}
	resp := &plugin.RunResponse{
		Hangup: false,
	}
	if json.Unmarshal([]byte(args), &arg) == nil {
		if arg.Timeout == 0 {
			arg.Timeout = 120
		}
		if arg.Provider == "list" {
			r := ListProviders()
			re, err := json.Marshal(r)
			if err != nil {
				return nil, err
			}
			resp.Response = string(re)
		} else {
			err := ExecProvider(ctx, arg.Provider, arg.Content)
			if err != nil {
				return nil, fmt.Errorf("failed to execute shell: %w", err)
			}
			resp.Response = "success"
		}
	} else {
		return nil, fmt.Errorf("invalid arguments")
	}
	return resp, nil
}

func (s *Service) Continue(ctx context.Context, session string) (*plugin.RunResponse, error) {
	return nil, nil
}

func (s *Service) Stop() error {
	return nil
}
