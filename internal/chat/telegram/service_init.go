package telegram

import (
	"context"

	"github.com/go-viper/mapstructure/v2"
)

// Init initial service
func (s *Service) Init(ctx context.Context, conf any) error {
	// get config
	var cfg Config
	err := mapstructure.Decode(conf, &cfg)
	s.cfg = cfg
	if err != nil {
		return err
	}
	return nil
}
