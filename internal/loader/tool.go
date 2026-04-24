package loader

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/tool"

	_ "github.com/manboster/manboster/internal/tool/datetime"
)

func LoadToolCallProviders(ctx context.Context) ([]tool.Provider, error) {
	toolCallProviders := tool.Providers()
	for _, provider := range toolCallProviders {
		color.Blue(fmt.Sprintf("[Manboster Loader] Loading tool call provider %q...", provider.DisplayName()))
		err := provider.Init(ctx)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Loader] We encountered an problem while loading tool call provider %q: %q", provider.DisplayName(), err))
			continue
		}
		go func() {
			err := provider.Start(ctx)
			if err != nil {
				color.Red(fmt.Sprintf("[Manboster Loader] We encountered an problem while polling tool call provider %q: %q", provider.DisplayName(), err))
			}
		}()

		// str, _ := json.MarshalIndent(provider.Args(), "", " ")
		// fmt.Printf("%s", string(str))
	}
	return toolCallProviders, nil
}
