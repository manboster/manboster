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
func (l *Loader) Load(ctx context.Context) error {
	color.Blue(fmt.Sprintf("[Manboster Loader] Validating Configuration..."))
	err := l.cfg.Validate()
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
	l.db = dbi
	repo := repository.New(dbi.Instance())
	l.repo = repo

	color.Blue(fmt.Sprintf("[Manboster Loader] Initializing LLM Providers..."))
	llmProviders, err := LoadLLMProviders(ctx, l.cfg.LLMs)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Loader] We encountered an error while initializing LLM Providers: %q", err))
		return err
	}
	if len(llmProviders) == 0 {
		color.Red(fmt.Sprintf("[Manboster Loader] We encountered an error while initializing LLM Providers: no llm provider available"))
		return fmt.Errorf("no llm provider available")
	}
	l.llmProviders = llmProviders
	// load default model
	l.loadDefaultModel(ctx)

	// load enabled tool call
	_, err = LoadToolCallProviders(ctx)

	// open a new engine
	color.Blue(fmt.Sprintf("[Manboster Loader] Initializing Manboster Engine..."))
	e, err := engine.New(l.cfg, repo, llmProviders)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Loader] We encountered an error while creating the engine: %q", err))
		return err
	}
	l.engine = e

	// load it, and enjoy it!
	err = e.Load(ctx)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Loader] We encountered an error while loading and running the engine: %q", err))
		return err
	}

	// we activate chats after loading engine
	err = l.LoadChats(ctx, l.cfg.Chats)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Loader] We encountered an error while loading and running the chat providers: %q", err))
		return err
	}

	return nil
}
