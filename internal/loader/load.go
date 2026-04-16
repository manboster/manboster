package loader

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/database"
	"github.com/manboster/manboster/internal/engine"
	"github.com/manboster/manboster/internal/repository"
)

// Load loads the loader
func Load(ctx context.Context) error {
	color.Blue(fmt.Sprintf("[Manboster Loader] Reading Configuration..."))
	cfg := config.Read()
	color.Blue(fmt.Sprintf("[Manboster Loader] Validating Configuration..."))
	err := cfg.Validate()
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Loader] We encountered an error while validating the configuration: %q", err))
		return err
	}

	// initialize database
	dbi := &database.Client{}
	dbPath := config.Read().App.DBPath
	// if there is no manboster.db definition, fallback to same folder
	if dbPath == "" {
		dbPath = "manboster.db"
	}
	color.Blue(fmt.Sprintf("[Manboster Loader] Initializing Manboster Database Repository..."))
	err = dbi.Init(dbPath)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Loader] We encountered an error while loading the database: %q", err))
		return err
	}
	repo := repository.New(dbi.Instance())
	color.Blue(fmt.Sprintf("[Manboster Loader] Initializing Manboster Engine..."))
	// open a new engine
	e, err := engine.New(cfg, repo)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Loader] We encountered an error while creating the engine: %q", err))
		return err
	}

	// load it, and enjoy it!
	err = e.Load(ctx)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Loader] We encountered an error while loading and running the engine: %q", err))
		return err
	}
	return nil
}
