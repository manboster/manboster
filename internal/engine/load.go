package engine

import (
	"context"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/database"
)

func (e *Engine) Load(ctx context.Context) error {
	// initialize variables

	// initialize database
	dbi := &database.Client{}
	// TODO: hard write manboster.db
	dbPath := config.Read().App.DBPath
	// if there is no manboster.db definition, fallback to same folder
	if dbPath == "" {
		dbPath = "manboster.db"
	}
	err := dbi.Init(dbPath)
	if err != nil {
		return err
	}

	// TODO: get model data from SQLite(Repository)
	// First, we activate LLMs.
	llmProviders, err := loadLLM(ctx, e.config.LLMs)
	if err != nil {
		return err
	}
	e.llmProviders = llmProviders

	// Then, we activate chats.
	err = e.loadChats(ctx)
	if err != nil {
		return err
	}
	return nil
}
