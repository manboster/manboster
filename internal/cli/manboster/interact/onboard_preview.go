package interact

import (
	"fmt"
	"strings"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/spec/cli"
)

func runOnboardPreview(p cli.Provider, c config.Config) (bool, error) {
	confDescription := strings.Builder{}
	confDescription.WriteString("# Before you proceed, you need to review what you have entered. \n")
	confDescription.WriteString("If anything is incorrect, please use Ctrl+C to quit and restart it with 'manboster onboard'.\n\n")

	confDescription.WriteString(fmt.Sprintf("You configured %d chat providers\n\n", len(c.Chats)))
	for i, _ := range c.Chats {
		confDescription.WriteString(fmt.Sprintf("#%d: %s's Configuration:\n\n %s\n\n", i+1, c.Chats[i].Provider, c.Chats[i].Configuration))
	}

	confDescription.WriteString(fmt.Sprintf("You configured %d llm providers\n\n", len(c.LLMs)))
	for i, _ := range c.LLMs {
		confDescription.WriteString(fmt.Sprintf("#%d's Configuration: \n\n%s \n\n", i+1, c.LLMs[i].Configuration))
	}

	confDescription.WriteString(fmt.Sprintf("You configured %d tool providers\n\n", len(c.Tools)))
	for i, _ := range c.Tools {
		confDescription.WriteString(fmt.Sprintf("#%d: %s's Configuration: \n\n", i+1, c.Tools[i].Name))
		if c.Tools[i].Configuration != nil {
			confDescription.WriteString(fmt.Sprintf("%s \n\n", c.Tools[i].Configuration))
		}
	}

	if c.Hachimi.Enabled {
		confDescription.WriteString(fmt.Sprintf("You enabled hachimi features\n\n"))
		for i, _ := range c.Hachimi.Hachimi {
			hcm := c.Hachimi.Hachimi[i]
			confDescription.WriteString(fmt.Sprintf("#%d: %s's Configuration: \n\n%s \n\n", i+1, hcm.Provider, hcm.Configuration))
		}
	} else {
		confDescription.WriteString(fmt.Sprintf("You disabled hachimi feature.\n\n"))
	}

	confDescription.WriteString("If there is no problem, you can continue writing the configuration.\n\n")
	confDesc := confDescription.String()

	return p.Prompt(confDesc, "Do you want to continue?", "Continue and write configuration", "There is something wrong")
}
