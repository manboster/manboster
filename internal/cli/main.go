package cli

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/database"
	"github.com/manboster/manboster/internal/engine"
	"github.com/manboster/manboster/internal/repository"
	"github.com/spf13/cobra"

	_ "github.com/manboster/manboster/internal/chat/telegram"
	_ "github.com/manboster/manboster/internal/llm/oai_compat"
	_ "github.com/manboster/manboster/internal/llm/openrouter"
)

// main is the entrypoint function that when user runs 'manboster'.
func main(cmd *cobra.Command, args []string) {
	// output welcome
	color.Cyan("Welcome to Manboster!")
	color.Blue("[Manboster Client] Your Lobster is on the way, please wait...")

	cfg := config.Read()
	err := cfg.Validate()
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	// initialize database
	dbi := &database.Client{}
	dbPath := config.Read().App.DBPath
	// if there is no manboster.db definition, fallback to same folder
	if dbPath == "" {
		dbPath = "manboster.db"
	}
	err = dbi.Init(dbPath)
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
	repo := repository.New(dbi.Instance())

	// create a universal context for this application
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	// open a new engine
	e, err := engine.New(cfg, repo)
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	// load it, and enjoy it!
	err = e.Load(ctx)
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	<-ctx.Done()
	color.Yellow("[Manboster Client] Your Manboster is going to sleep, thank you for playing with it!")
}
