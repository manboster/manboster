package interact

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/spec/cli"
)

func runConfig(p cli.Provider, cfg config.Config) (config.Config, error) {
	err := p.Alert("Manboster Configuration Manager", "Work in progress...")
	if err != nil {
		return cfg, err
	}
	
	return cfg, nil
}
