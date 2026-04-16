package telegram

import (
	"context"

	"github.com/go-viper/mapstructure/v2"
)

// Init initial service
func (s *Service) Init(ctx context.Context, conf any) error {
	// get config
	err := mapstructure.Decode(conf, s.cfg)
	if err != nil {
		return err
	}

	err = s.cfg.Validate()
	if err != nil {
		return err
	}

	return nil
}
