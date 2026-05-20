package openrouter

import (
	"context"

	"github.com/go-viper/mapstructure/v2"
	"github.com/manboster/manboster/internal/llm/oai_compat"
)

func (s *Service) Init(ctx context.Context, config any) error {
	// read config
	var conf Config
	err := mapstructure.Decode(config, &conf)
	if err != nil {
		return err
	}

	oaiInstance := &oai_compat.Service{}
	err = oaiInstance.InitWithConfig(ctx, conf.Config)
	if err != nil {
		return err
	}
	s.oaiInstance = oaiInstance

	return nil
}
