package datetime

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func (s *Service) Init(ctx context.Context) error {
	return nil
}

func (s *Service) Start(ctx context.Context) error {
	return nil
}

func (s *Service) Run(ctx context.Context, args string) (string, error) {
	arg := RunArgs{}
	if json.Unmarshal([]byte(args), &arg) != nil {
		switch arg.Name {
		case "date":
			return time.Now().Format("2006-01-02"), nil
		case "time":
			return time.Now().Format("15:04:05"), nil
		default:
			return "", fmt.Errorf("unknown argument %q", arg.Name)
		}
	} else {
		return "", fmt.Errorf("invalid arguments")
	}
}

func (s *Service) Close() error {
	return nil
}
