package openrouter

import (
	"context"
	"strconv"

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

	// convert openrouter config to oai-compat one
	var oaiConfig oai_compat.Config
	oaiConfig.ProviderName = "openrouter"
	oaiConfig.ProviderDisplayName = "OpenRouter"
	oaiConfig.Model = conf.Model
	oaiConfig.ApiKey = conf.ApiKey
	oaiConfig.BaseURL = "https://openrouter.ai/api/v1" // fixed openrouter api calls

	if conf.ID > 0 {
		oaiConfig.ProviderName += "-" + strconv.Itoa(conf.ID)
		oaiConfig.ProviderDisplayName += "-" + strconv.Itoa(conf.ID)
	}

	oaiInstance := &oai_compat.Service{}
	err = oaiInstance.InitWithConfig(ctx, oaiConfig)
	if err != nil {
		return err
	}
	s.oaiInstance = oaiInstance

	return nil
}
