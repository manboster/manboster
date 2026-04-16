package telegram

import (
	"context"

	"github.com/go-viper/mapstructure/v2"
)

// Init initial service
func (s *Service) Init(ctx context.Context, conf any) error {
	var cfg Config
	// get config
	err := mapstructure.Decode(conf, &cfg)
	if err != nil {
		return err
	}

	s.cfg = &cfg
	err = s.cfg.Validate()
	if err != nil {
		return err
	}

	return nil
}
