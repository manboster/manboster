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
		color.Red(fmt.Sprintf(i18n.T(keys.DaemonStartError), err))
		return
	}

	if d != nil {
		color.Green(i18n.T(keys.DaemonStartSuccess))
		return
	}

	defer func(ctx *daemon.Context) {
		err := ctx.Release()
		if err != nil {
			color.Red(fmt.Sprintf(i18n.T(keys.DaemonStopError), err))
		}
	}(ctx.DaemonCtx)

	err = config.Init()
	if errors.Is(err, config.ErrNoConfig) {
		color.Red(i18n.T(keys.DaemonNoConfig))
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
		color.Red(fmt.Sprintf(i18n.T(keys.DaemonStatusError), err))
		return
	}
	if d != nil {
		err = d.Signal(syscall.SIGTERM)
		if err != nil {
			color.Red(fmt.Sprintf(i18n.T(keys.DaemonStopError), err))
			return
		}
		color.Green(i18n.T(keys.DaemonStopSuccess))
	} else {
		color.Yellow(i18n.T(keys.DaemonStopStopped))
	}
}

// restartCommandExecutor restarts the daemon
func restartCommandExecutor(cmd *cobra.Command, args []string) {
	color.Cyan(i18n.T(keys.DaemonRestartMessage))
	stopCommandExecutor(cmd, args)
	time.Sleep(3 * time.Second)
	startCommandExecutor(cmd, args)
}

// statusCommandExecutor returns current daemon status of Manboster
func statusCommandExecutor(cmd *cobra.Command, args []string) {
	d, err := ctx.DaemonCtx.Search()
	if err != nil {
		color.Red(fmt.Sprintf(i18n.T(keys.DaemonStatusError), err))
		return
	}
	if d != nil {
		err = d.Signal(syscall.Signal(0))
		if err != nil {
			color.Red(fmt.Sprintf(i18n.T(keys.DaemonStatusNotRunning), config.Path("manboster.pid"), err))
			return
		}
		color.Green(fmt.Sprintf(i18n.T(keys.DaemonStatusRunning), config.Path("manboster.log")))
	} else {
		color.Red(i18n.T(keys.DaemonStopStopped))
	}
}
