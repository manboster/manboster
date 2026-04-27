package datetime

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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
		switch arg.Name {
		case "date":
			resp.Response = time.Now().Format("2006-01-02")
		case "time":
			resp.Response = time.Now().Format("15:04:05")
		default:
			return nil, fmt.Errorf("unknown argument %q", arg.Name)
		}
	} else {
		return nil, fmt.Errorf("invalid arguments")
	}
	return resp, nil
}

func (s *Service) Continue(ctx context.Context, session string) (*plugin.RunResponse, error) {
	return nil, nil
}

func (s *Service) Close() error {
	return nil
}
