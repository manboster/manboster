package interact

import (
	"fmt"
	"strings"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/spec/cli"
)

func runOnboardPreview(p cli.Provider, c config.Config) (bool, error) {
	confDescription := strings.Builder{}
	confDescription.WriteString(i18n.T(keys.OnboardPreviewTitle) + "\n")
	confDescription.WriteString(i18n.T(keys.OnboardPreviewRestart) + "\n\n")

	confDescription.WriteString(fmt.Sprintf(i18n.T(keys.OnboardPreviewChatCount), len(c.Chats)) + "\n\n")
	for i, _ := range c.Chats {
		confDescription.WriteString(fmt.Sprintf("#%d: %s's Configuration:\n\n %s\n\n", i+1, c.Chats[i].Provider, c.Chats[i].Configuration))
	}

	confDescription.WriteString(fmt.Sprintf(i18n.T(keys.OnboardPreviewLLMCount), len(c.LLMs)) + "\n\n")
	for i, _ := range c.LLMs {
		confDescription.WriteString(fmt.Sprintf("#%d's Configuration: \n\n%s \n\n", i+1, c.LLMs[i].Configuration))
	}

	confDescription.WriteString(fmt.Sprintf(i18n.T(keys.OnboardPreviewToolCount), len(c.Tools)) + "\n\n")
	for i, _ := range c.Tools {
		confDescription.WriteString(fmt.Sprintf("#%d: %s's Configuration: \n\n", i+1, c.Tools[i].Name))
		if c.Tools[i].Configuration != nil {
			confDescription.WriteString(fmt.Sprintf("%s \n\n", c.Tools[i].Configuration))
		}
	}

	if c.Hachimi.Enabled {
		confDescription.WriteString(i18n.T(keys.OnboardPreviewHachimiEnabled) + "\n\n")
		for i, _ := range c.Hachimi.Hachimi {
			hcm := c.Hachimi.Hachimi[i]
			confDescription.WriteString(fmt.Sprintf("#%d: %s's Configuration: \n\n%s \n\n", i+1, hcm.Provider, hcm.Configuration))
		}
	} else {
		confDescription.WriteString(i18n.T(keys.OnboardPreviewHachimiDisabled) + "\n\n")
	}

	confDescription.WriteString(i18n.T(keys.OnboardPreviewContinue) + "\n\n")
	confDesc := confDescription.String()

	return p.Prompt(confDesc, i18n.T(keys.OnboardPreviewConfirm), "Continue and write configuration", i18n.T(keys.OnboardPreviewProblem))
}
