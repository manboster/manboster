package interact

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/spec/cli"
)

func runAppConfig(p cli.Provider, cfg config.Config) (config.AppConfig, error) {
	return runOnboardAPPConfig(p, cfg)
}
