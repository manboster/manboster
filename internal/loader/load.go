package loader

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/database"
	"github.com/manboster/manboster/internal/engine"
	"github.com/manboster/manboster/internal/repository"
	chatType "github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
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
	database.DBInstance = dbi
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
	// load default model
	llmProvidersMap := make(map[string]llm.Provider)
	for _, p := range llmProviders {
		llmProvidersMap[p.Name()] = p
	}
	l.llmProviders = llmProvidersMap
	l.loadDefaultModel(ctx)

	// load enabled tool call
	color.Blue("[Manboster Loader] Loaded Tool Call Providers...")
	tool, err := LoadToolCallProviders(ctx, l.cfg)
	if err != nil {
		color.Yellow(fmt.Sprintf("[Manboster Loader] Failed to load Tool Call Providers: %q", err))
	}

	hachimiLoaded := false

	// check hachimi enable status and open based on hachimi's provider name
	if l.cfg.Hachimi.Enabled {
		color.Blue(fmt.Sprintf("[Manboster Loader] Hachimi enabled, initializing Hachimi Providers..."))
		for _, conf := range l.cfg.Hachimi.Hachimi {
			if l.cfg.Hachimi.Provider == conf.Provider {
				color.Blue(fmt.Sprintf(fmt.Sprintf("[Manboster Loader] Initializing Hachimi Provider %q...", conf.Provider)))
				hProvider, err := LoadHachimiProvider(ctx, conf)
				if err != nil {
					color.Yellow(fmt.Sprintf("[Manboster Loader] Could not load Hachimi Provider: %q", err))
					break
				}
				l.hachimiProvider = hProvider
				hachimiLoaded = true
			}
		}
	}

	// we activate chats after loading engine
	cProviders, err := l.LoadChats(ctx, l.cfg.Chats)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Loader] We encountered an error while loading and running the chat providers: %q", err))
		return err
	}
	l.chatProviders = cProviders

	chatProvidersMap := make(map[string]chatType.Provider)
	for _, p := range cProviders {
		chatProvidersMap[p.Name()] = p
	}

	// open a new engine
	color.Blue(fmt.Sprintf("[Manboster Loader] Initializing Manboster Engine..."))

	if _, err := os.Stat(config.Path("SOUL.md")); os.IsNotExist(err) {
		color.Cyan(fmt.Sprintf("[Manboster] Tips: You can create file %q to define your Lobster's soul!", config.Path("SOUL.md")))
	}

	e, err := engine.New(l.cfg, repo, llmProvidersMap, chatProvidersMap, tool, l.hachimiProvider, &hachimiLoaded)
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

	return nil
}
