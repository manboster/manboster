package interact

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/spec/cli"
)

func OnboardWarningPrompt(provider cli.Provider) (bool, error) {
	t, err := provider.Prompt(
		i18n.T(keys.OnboardWarningRiskTitle),
		i18n.T(keys.OnboardWarningRiskPrompt),
		i18n.T(keys.OnboardWarningAccept),
		i18n.T(keys.OnboardWarningExit),
	)
	if err != nil || !t {
		return t, err
	}

	if config.VersionType(config.CurrentChannel) != config.ChannelStable {
		return provider.Prompt(
			i18n.T(keys.OnboardWarningUnstableTitle),
			i18n.T(keys.OnboardWarningUnstablePrompt),
			i18n.T(keys.OnboardWarningAccept),
			i18n.T(keys.OnboardWarningExit),
		)
	}

	return true, nil
}
