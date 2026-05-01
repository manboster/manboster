package browser

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
	if !s.isReady {
		return nil, fmt.Errorf("the browser is not ready, please wait for a while or check out whether there is an error or not")
	}

	sessID, ok := ctx.Value("session_id").(string)
	if !ok {
		return nil, fmt.Errorf("session_id not found in context")
	}

	if json.Unmarshal([]byte(args), &arg) == nil {
		switch arg.Name {
		case NameTypeWebpage:
			res, err := s.ScrapWebpage(ctx, arg.URL, arg.ScrapType, arg.ResponseType, sessID)
			if err != nil {
				return nil, err
			}
			resp.Response = res
		case NameTypeSearch:
			res, err := s.doWebSearch(ctx, arg.Keywords, arg.Engine, ResponseTypeRaw, sessID)
			if err != nil {
				return nil, err
			}
			resp.Response = res
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
