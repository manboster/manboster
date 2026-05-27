package daemon

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/cli/manboster/app"
	"github.com/manboster/manboster/internal/cli/manboster/ctx"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"
)

// startCommandExecutor starts daemon
func startCommandExecutor(cmd *cobra.Command, args []string) {
	d, err := ctx.DaemonCtx.Reborn()
	if err != nil {
		color.Red(fmt.Sprintf(i18n.T(keys.AppDaemonStartError), err))
		return
	}

	if d != nil {
		color.Green(i18n.T(keys.AppDaemonStartSuccess))
		return
	}

	defer func(ctx *daemon.Context) {
		err := ctx.Release()
		if err != nil {
			color.Red(fmt.Sprintf(i18n.T(keys.AppDaemonStopError), err))
		}
	}(ctx.DaemonCtx)

	err = config.Init()
	if errors.Is(err, config.ErrNoConfig) {
		color.Red(i18n.T(keys.AppDaemonNoConfig))
		os.Exit(0)
	} else if err != nil {
		panic(err)
	}
	app.MainInner()
}

// stopCommandExecutor stops daemon
func stopCommandExecutor(cmd *cobra.Command, args []string) {
	d, err := ctx.DaemonCtx.Search()
	if err != nil {
		color.Red(fmt.Sprintf(i18n.T(keys.AppDaemonStatusError), err))
		return
	}
	if d != nil {
		err = d.Signal(syscall.SIGTERM)
		if err != nil {
			color.Red(fmt.Sprintf(i18n.T(keys.AppDaemonStopError), err))
			return
		}
		color.Green(i18n.T(keys.AppDaemonStopSuccess))
	} else {
		color.Yellow(i18n.T(keys.AppDaemonStopStopped))
	}
}

// restartCommandExecutor restarts the daemon
func restartCommandExecutor(cmd *cobra.Command, args []string) {
	color.Cyan(i18n.T(keys.AppDaemonRestartMessage))
	stopCommandExecutor(cmd, args)
	time.Sleep(3 * time.Second)
	startCommandExecutor(cmd, args)
}

// statusCommandExecutor returns current daemon status of Manboster
func statusCommandExecutor(cmd *cobra.Command, args []string) {
	d, err := ctx.DaemonCtx.Search()
	if err != nil {
		color.Red(fmt.Sprintf(i18n.T(keys.AppDaemonStatusError), err))
		return
	}
	if d != nil {
		err = d.Signal(syscall.Signal(0))
		if err != nil {
			color.Red(fmt.Sprintf(i18n.T(keys.AppDaemonStatusNotRunning), config.Path("manboster.pid"), err))
			return
		}
		color.Green(fmt.Sprintf(i18n.T(keys.AppDaemonStatusRunning), config.Path("manboster.log")))
	} else {
		color.Red(i18n.T(keys.AppDaemonStopStopped))
	}
}
