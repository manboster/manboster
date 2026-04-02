package oai_compat

import (
	"context"
	"net/http"

	"github.com/go-viper/mapstructure/v2"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
	"github.com/sashabaranov/go-openai"
)

func (s *Service) Init(ctx context.Context, config any) error {
	// read config
	var conf Config
	err := mapstructure.Decode(config, &conf)
	if err != nil {
		return err
	}
	return s.InitWithConfig(ctx, conf)
}

func (s *Service) InitWithConfig(ctx context.Context, config Config) error {
	s.cfg = config

	// set default headers, then overwrite it with client configuration.
	defaultHeaders := make(map[string]string)
	defaultHeaders["HTTP-Referer"] = "https://github.com/manboster/manboster"
	defaultHeaders["X-Client"] = "Manboster"

	// overwrite it with client
	if len(config.Headers) != 0 {
		for key, value := range config.Headers {
			defaultHeaders[key] = value
		}
	}

	// create a brand-new configuration using configs.
	oaiConfig := openai.DefaultConfig(config.ApiKey)
	oaiConfig.BaseURL = config.BaseURL

	// profiling headers to each request.
	httpClient := &http.Client{
		Transport: &headerTransport{
			base:    http.DefaultTransport,
			headers: defaultHeaders,
		},
	}
	oaiConfig.HTTPClient = httpClient

	oaiCli := openai.NewClientWithConfig(oaiConfig)
	s.cli = oaiCli
	return nil
}

func init() {
	llm.Register("oai-compat", func() llm.Provider {
		return &Service{}
	})
	config.Register("llm:oai-compat", func() config.Provider {
		return &Config{}
	})
}
