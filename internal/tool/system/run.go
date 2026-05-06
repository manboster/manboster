package system

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
		switch arg.Name {
		case NameOSInfo:
			sys, err := getSystemInfo(ctx)
			if err != nil {
				return resp, fmt.Errorf("failed to get system info: %w", err)
			}
			jsonify, err := json.Marshal(sys)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal system info: %w", err)
			}
			resp.Response = string(jsonify)
		case NameProcess:
			switch arg.Action {
			case ActionList:
				sys, err := listProcesses(ctx)
				if err != nil {
					return resp, fmt.Errorf("failed to get process list: %w", err)
				}
				jsonify, err := json.Marshal(sys)
				if err != nil {
					return nil, fmt.Errorf("failed to marshal process: %w", err)
				}
				resp.Response = string(jsonify)
			case ActionInfo:
				sys, err := getProcessInfo(ctx, int32(arg.PID))
				if err != nil {
					return resp, fmt.Errorf("failed to get process list: %w", err)
				}
				jsonify, err := json.Marshal(sys)
				if err != nil {
					return nil, fmt.Errorf("failed to marshal process: %w", err)
				}
				resp.Response = string(jsonify)
			case ActionKill:
				err := killProcess(ctx, int32(arg.PID))
				if err != nil {
					return resp, fmt.Errorf("failed to kill process: %w", err)
				}
				resp.Response = "OK"
			default:
				return resp, fmt.Errorf("unknown action: %s", arg.Action)
			}
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

func (s *Service) Stop() error {
	return nil
}
