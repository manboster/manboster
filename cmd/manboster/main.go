package main

import (
	"errors"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/cli"
	"github.com/manboster/manboster/internal/config"
	"github.com/spf13/cobra"
)

// Manboster: Your Personal Manbo Lobster!
// Powered by chihuo2104 (c) 2026.
// Last Update: 2026.4.14

func main() {
	err := config.Init()
	if errors.Is(err, config.ErrNoConfig) {
		cli.ConfigCmdRun(&cobra.Command{}, os.Args[1:])
		color.Green("Successfully created config.yaml, open Manboster again and enjoy it!")
		time.Sleep(1 * time.Second)
		os.Exit(0)
	} else if err != nil {
		panic(err)
	}

	cli.Init()
}
