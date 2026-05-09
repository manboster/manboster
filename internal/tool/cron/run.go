package cron

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/manboster/manboster/spec/plugin"
)

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
	userId, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, fmt.Errorf("user_id not found in context")
	}

	if json.Unmarshal([]byte(args), &arg) == nil {
		switch arg.Name {
		case NameSet:
			if arg.MessageType != MessagePrompt && arg.MessageType != MessageText {
				return nil, fmt.Errorf("invalid argument '%s' in messageType", arg.Name)
			}
			err := s.Create(ctx, arg, chatID, chatProvider, userId)
			if err != nil {
				return nil, err
			}
			resp.Response = "success"
		case NameGet:
			data, err := s.Get(ctx, arg.JobName)
			if err != nil {
				return nil, err
			}
			jsonify, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}
			resp.Response = string(jsonify)
		case NameList:
			list, err := s.List(ctx, chatProvider, chatID)
			if err != nil {
				return nil, err
			}
			jsonify, err := json.Marshal(list)
			if err != nil {
				return nil, err
			}
			resp.Response = string(jsonify)
		case NameDelete:
			err := s.Delete(ctx, arg.JobName)
			if err != nil {
				return nil, err
			}
			resp.Response = "success"
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
