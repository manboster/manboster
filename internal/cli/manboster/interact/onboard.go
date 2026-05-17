package interact

import (
	"fmt"

	_ "github.com/manboster/manboster/internal/chat/all"
	"github.com/manboster/manboster/internal/config"
	_ "github.com/manboster/manboster/internal/hachimi/all"
	_ "github.com/manboster/manboster/internal/llm/all"
	_ "github.com/manboster/manboster/internal/tool/all"
	"github.com/manboster/manboster/spec/cli"
)

func runOnboardConfig(p cli.Provider) (config.Config, error) {
	conf := config.Config{}

	allowed, err := OnboardWarningPrompt(p)
	if err != nil {
		return conf, err
	}
	if !allowed {
		return conf, fmt.Errorf("you rejected the warning, in order to protect you, we skip your installation progress")
	}

	err = p.Alert("Manboster Configuration Wizard", "Welcome to the Manboster Configuration Wizard. Enjoy your experience with your little Manbo!")
	if err != nil {
		return conf, err
	}

	return conf, nil
}
