package interact

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/tool"
	_ "github.com/manboster/manboster/internal/tool/all"
	"github.com/manboster/manboster/spec/cli"
)

func runOnboardToolConfig(p cli.Provider) ([]config.ToolConfig, error) {
	var conf []config.ToolConfig
	var toolProviders []tool.Provider

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	toolAvails := tool.AvailProviders()
	for _, t := range toolAvails {
		provider, err := tool.GetProvider(t)
		if err != nil {
			continue
		}

		toolProviders = append(toolProviders, provider)
	}
	options := cli.BuildOptionsWithMetadata[tool.Provider](toolProviders, nil)
	selects, err := p.MultiSelect("Select and activate the tools you want to use", "Please select the tool you want to use, please be careful to select as they will be the tool call of AIs.", options, nil, func(options []cli.Option) error {
		if options == nil {
			return nil
		}
		for _, o := range options {
			mark := false
			for _, tp := range toolProviders {
				if o.Value == tp.Name() {
					mark = true
					break
				}
			}
			if !mark {
				return fmt.Errorf("%s is not a valid tool type", o.Value)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	for _, provider := range toolProviders {
		for i, selected := range selects {
			if selected.Value == provider.Name() {
				if provider.Config() != nil {
					createConfig, err := CreateConfig(ctx, p, provider.Config())
					if err != nil {
						return nil, err
					}
					conf = append(conf, config.ToolConfig{
						Name:          provider.Name(),
						Configuration: createConfig,
					})
				} else {
					conf = append(conf, config.ToolConfig{
						Name:          provider.Name(),
						Configuration: nil,
					})
				}

				selects = append(selects[:i], selects[i+1:]...)
				break
			}
		}
	}

	return conf, nil
}
