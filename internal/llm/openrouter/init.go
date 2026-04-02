package openrouter

import (
	"context"

	"github.com/go-viper/mapstructure/v2"
	"github.com/manboster/manboster/internal/llm"
	oai_compat2 "github.com/manboster/manboster/internal/llm/oai_compat"
)

func (s *Service) Init(ctx context.Context, config any) error {
	// read config
	var conf oai_compat2.Config
	err := mapstructure.Decode(config, &conf)
	if err != nil {
		return err
	}
	conf.BaseURL = "https://openrouter.ai/api/v1" // fixed openrouter api calls
	oaiInstance := &oai_compat2.Service{}
	err = oaiInstance.InitWithConfig(ctx, conf)
	if err != nil {
		return err
	}
	s.oaiInstance = oaiInstance

	return nil
}

func init() {
	llm.Register("openrouter", func() llm.Provider {
		return &Service{}
	})
}
