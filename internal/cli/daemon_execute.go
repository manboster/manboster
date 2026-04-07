package cli

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"
)

var ctx = &daemon.Context{
	PidFileName: config.Path("manboster.pid"),
	PidFilePerm: 0644,
	LogFileName: config.Path("manboster.log"),
	LogFilePerm: 0640,
	WorkDir:     config.Path(""),
	Umask:       027,
}

// startCommandExecutor starts daemon
func startCommandExecutor(cmd *cobra.Command, args []string) {
	d, err := ctx.Reborn()
	if err != nil {
		color.Red(fmt.Sprintf("An error has been occurred when starting Manboster daemon, error: %v\n Please be sure that this daemon is not running.", err))
		return
	}

	if d != nil {
		color.Green("Manboster daemon started, you can chat with your Lobster after a while!")
		return
	}

	defer func(ctx *daemon.Context) {
		err := ctx.Release()
		if err != nil {
			color.Red(fmt.Sprintf("An error has been occurred when stopping Manboster daemon, error: %v", err))
		}
	}(ctx)

	err = config.Init()
	if errors.Is(err, config.ErrNoConfig) {
		color.Red("There is no configuration file available, please run at least once!")
		os.Exit(0)
	} else if err != nil {
		panic(err)
	}
	main(cmd, args)
}

// stopCommandExecutor stops daemon
func stopCommandExecutor(cmd *cobra.Command, args []string) {
	d, err := ctx.Search()
	if err != nil {
		color.Red(fmt.Sprintf("An error has been occurred when stopping Manboster daemon, error: %v\n Please be sure that you have started this daemon.", err))
		return
	}

	// stop the daemon
	err = d.Signal(syscall.SIGTERM)
	if err != nil {
		color.Red(fmt.Sprintf("An error has been occurred when stopping Manboster daemon, error: %v", err))
		return
	}
	color.Green("Manboster daemon stopped, thank you for playing with your Lobster!")
}

// restartCommandExecutor restarts the daemon
func restartCommandExecutor(cmd *cobra.Command, args []string) {
	color.Cyan("Restarting Manboster daemon, please wait...")
	stopCommandExecutor(cmd, args)
	time.Sleep(3 * time.Second)
	startCommandExecutor(cmd, args)
}

// statusCommandExecutor returns current daemon status of Manboster
func statusCommandExecutor(cmd *cobra.Command, args []string) {
	d, err := ctx.Search()
	if err != nil {
		color.Red(fmt.Sprintf("An error has been occurred when getting Manboster daemon PID file, error: %v\nMaybe the daemon is not running.", err))
		return
	}
	// get running status
	err = d.Signal(syscall.Signal(0))
	if err != nil {
		color.Red(fmt.Sprintf("Manboster is not running currently, please delete PID files in %s to reset daemon status or run 'manboster reset' , error: %v", config.Path("manboster.pid"), err))
		return
	}

	color.Green(fmt.Sprintf("Manboster daemon is running, you can view the log data in %s", config.Path("manboster.log")))
}
