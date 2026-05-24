package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/cli/manboster/ctx"
	"github.com/manboster/manboster/internal/cli/manboster/interact"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/loader"
	"github.com/manboster/manboster/internal/tool"
	"github.com/spf13/cobra"

	_ "github.com/manboster/manboster/internal/chat/all"
	_ "github.com/manboster/manboster/internal/llm/all"
)

// Main is the entrypoint function that when user runs 'manboster'.
func Main(cmd *cobra.Command, args []string) {
	color.Cyan(i18n.T(keys.AppWelcome))
	color.Blue(i18n.T(keys.AppLoading))

	_, err := ctx.DaemonCtx.Search()
	if err == nil {
		color.Red(i18n.T(keys.AppDaemonRunning))
		color.Red(i18n.T(keys.AppDaemonRunningQuit))
		os.Exit(1)
	}

	MainInner()
}

func MainInner() {
	err := config.Init()
	if errors.Is(err, config.ErrNoConfig) {
		color.Yellow(i18n.T(keys.AppConfigNotFound))
		interact.OnboardConfigCmdRun(&cobra.Command{}, os.Args[1:])
		color.Green(i18n.T(keys.AppConfigCreated))
		time.Sleep(1 * time.Second)
		os.Exit(0)
	} else if err != nil {
		color.Red(fmt.Sprintf(i18n.T(keys.AppConfigInitError), err))
	}

	// create a universal context for this application
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	color.Blue(i18n.T(keys.AppReadingConfig))
	tool.IsLoading = true
	loaderInstance := loader.New(new(config.Read()))
	err = loaderInstance.Load(ctx)
	if err != nil {
		color.Red(fmt.Sprintf(i18n.T(keys.AppLoaderError), err))
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}

	<-ctx.Done()
	color.Yellow(i18n.T(keys.AppGoodbye))
	time.Sleep(1 * time.Second)
}
