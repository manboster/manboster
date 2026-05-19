package interact

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/spec/cli"
)

func runConfigWrite(p cli.Provider, cfg config.Config) error {
	confPath := config.Path("config.yaml")
	if _, err := os.Stat(confPath); os.IsExist(err) {
		confirm, err := p.Prompt(`
There is already a configuration file in your workspace!
If you continue, it will override the original configuration.
If you`, "Do you want to continue?", "Continue", "Save to other directory")
		if err != nil {
			return err
		}

		if !confirm {
			cp, err := p.Input("Configuration write path", "Please give the path you want to write this configuration file to", "", false, func(input string) error {
				if input == "" {
					return fmt.Errorf("no path specified")
				}
				return nil
			})
			if err != nil {
				return err
			}

			confPath = fmt.Sprintf("%s", cp)
			confPath = filepath.Join(confPath, "config.yaml")
		}
	}

	err := config.Write(cfg, confPath)
	if err != nil {
		confErr := p.Alert("Manboster Configuration Wizard", fmt.Sprintf("Failed to write configuration file: %s", err))
		if confErr != nil {
			return confErr
		}
		return err
	}

	err = p.Alert("Manboster Configuration Wizard", fmt.Sprintf("Successfully wrote configuration file to %s!\nEnjoy your little Manbo!", confPath))
	if err != nil {
		return err
	}
	return nil
}
