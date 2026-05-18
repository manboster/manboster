package interact

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/spec/cli"
)

// runOnboardAPPConfig runs app Config
func runOnboardAPPConfig(p cli.Provider, cfg config.Config) (config.AppConfig, error) {
	conf := config.AppConfig{}
	return conf, nil
}
