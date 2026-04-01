package openrouter

import (
	"context"

	"github.com/go-viper/mapstructure/v2"
	"github.com/sashabaranov/go-openai"
)

func (s *Service) Init(ctx context.Context, config any) error {
	// read config
	var conf Config
	err := mapstructure.Decode(config, &conf)
	if err != nil {
		return err
	}
	s.cfg = conf

	// create a brand new configuration using configs.
	orConfig := openai.DefaultConfig(conf.ApiKey)
	orConfig.BaseURL = "https://openrouter.ai/api/v1" // fixed openrouter api calls
	orCli := openai.NewClientWithConfig(orConfig)
	s.cli = orCli

	return nil
}
