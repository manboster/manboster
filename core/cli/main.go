package cli

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatih/color"
	"github.com/manboster/manboster/core/chat"
	"github.com/manboster/manboster/core/chat/providers/telegram"
	"github.com/manboster/manboster/core/config"
	"github.com/spf13/cobra"
)

// main is the entrypoint function that when user runs 'manboster'.
func main(cmd *cobra.Command, args []string) {
	// output welcome
	color.Cyan("Welcome to Manboster!")
	color.Blue("Your Lobster is on the way, please wait...")

	cfg := config.Read()

	current := int16(0)
	// check version
	if cfg.Version > current {
		color.Yellow("Configuration contains an unsupported version, if you want to use this configuration, please download the latest version. Or you can reconfigure it with `manboster config`.")
		os.Exit(1)
	}
	if cfg.Version < current {
		color.Yellow("Outdated configuration, if you want to use this configuration, please run `manboster upgrade` to upgrade your old data. Or you can reconfigure it with `manboster config`.")
		os.Exit(1)
	}

	// check valid configuration
	if len(cfg.Chats) == 0 {
		color.Red("Missing chat configuration, please reconfigure it with `manboster config`.")
		os.Exit(1)
	}
	if len(cfg.LLMs) == 0 {
		color.Red("Missing LLM configuration, please reconfigure it with `manboster config`.")
		os.Exit(1)
	}

	// create a universal context for this application
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	// TODO: get model data from SQLite(Repository)
	// but in MVP, we just choose the first one is enough.
	switch cfg.Chats[0].Provider {
	case "telegram":
		tgInstance := &telegram.Service{}
		err := tgInstance.Start(ctx, cfg.Chats[0].Configuration, func(message *chat.Message) {
			color.Blue("Got an message from Telegram by %s", message.Username)
			err := tgInstance.SendMessage(ctx, message)
			if err != nil {
				color.Red(err.Error())
				return
			}
		})
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
	}

	<-ctx.Done()
	color.Red("Your Manboster is going to sleep, thank you for playing with it!")
}
