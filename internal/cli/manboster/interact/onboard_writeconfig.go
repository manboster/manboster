package interact

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/spec/cli"
)

func runConfigWrite(p cli.Provider, cfg config.Config) error {
	confPath := config.Path("config.yaml")
	if _, err := os.Stat(confPath); os.IsExist(err) {
		confirm, err := p.Prompt(i18n.T(keys.OnboardWriteExisting), i18n.T(keys.OnboardWriteConfirm), "Continue", "Save to other directory")
		if err != nil {
			return err
		}

		if !confirm {
			cp, err := p.Input(i18n.T(keys.OnboardWritePathPrompt), i18n.T(keys.OnboardWritePathHelp), "", false, func(input string) error {
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
		confErr := p.Alert(i18n.T(keys.CliWizardTitle), i18n.Te(keys.OnboardWriteError, "", err))
		if confErr != nil {
			return confErr
		}
		return err
	}

	err = p.Alert(i18n.T(keys.CliWizardTitle), i18n.Te(keys.OnboardWriteSuccess, confPath, nil))
	if err != nil {
		return err
	}
	return nil
}
