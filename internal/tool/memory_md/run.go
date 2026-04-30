package memory_md

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/spec/plugin"
)

func (s *Service) Init(ctx context.Context, cfg any) error {
	err := os.MkdirAll(config.Path("memory"), 0755)
	if err != nil {
		return err
	}
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

	chatID, ok := ctx.Value("chat_id").(string)
	if !ok {
		return nil, fmt.Errorf("chat_id not found in context")
	}
	chatProvider, ok := ctx.Value("chat_provider").(string)
	if !ok {
		return nil, fmt.Errorf("chat_provider not found in context")
	}
	path := config.Path(filepath.Join("memory", fmt.Sprintf("memory-%s-%s.md", chatProvider, chatID)))

	if json.Unmarshal([]byte(args), &arg) == nil {
		switch arg.Name {
		case "get":
			data, err := os.ReadFile(path)
			if err != nil {
				return nil, fmt.Errorf("failed to read file %s: %w", path, err)
			}
			resp.Response = string(data)
		case "set":
			if len([]byte(arg.Value)) > 65536 {
				return nil, fmt.Errorf("failed to write: markdown entity too large")
			}
			err := os.WriteFile(path, []byte(arg.Value), 0644)
			if err != nil {
				return nil, fmt.Errorf("failed to write file %s: %w", path, err)
			}
			resp.Response = "Success"
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
