package interact

import (
	"context"
	"fmt"
	"strings"

	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/spec/cli"
)

type databaseSoulPageAction string

const (
	databaseSoulPageEdit   databaseSoulPageAction = _EDIT_
	databaseSoulPageDelete databaseSoulPageAction = _DELETE_
	databaseSoulPageQuit   databaseSoulPageAction = _QUIT_
)

func (a databaseSoulPageAction) Name() string { return string(a) }
func (a databaseSoulPageAction) DisplayName() string {
	switch a {
	case databaseSoulPageEdit:
		return "Edit this Soul's content"
	case databaseSoulPageDelete:
		return "Delete this Soul"
	case databaseSoulPageQuit:
		return "Quit"
	default:
		return ""
	}
}

func runDatabaseSoulConfig(p cli.Provider, repo repository.Repository) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var option cli.Option
	for {
		// reload on every iteration so changes are reflected
		souls, err := repo.GetAllSouls(ctx)
		if err != nil {
			return err
		}

		options := []cli.Option{createOption}
		for _, soul := range souls {
			options = append(options, cli.Option{
				Key:   fmt.Sprintf("`%s` (scope: %s)", soul.Name, strings.Join(soul.Scope, ", ")),
				Value: soul.Name,
			})
		}
		options = append(options, quitOption)

		summary := fmt.Sprintf("%d souls loaded.", len(souls))

		option, err = p.Select("Select a soul to manage.", summary, options, option.Value, func(o cli.Option) error {
			return nil
		})
		if err != nil {
			return err
		}

		if option.Value == _QUIT_ {
			return nil
		}

		if option.Value == _CREATE_ {
			if err := runDatabaseSoulCreate(ctx, p, repo); err != nil {
				return err
			}
			continue
		}

		var selectedSoul types.Soul
		selectedIndex := -1
		for i, soul := range souls {
			if soul.Name == option.Value {
				selectedSoul = soul
				selectedIndex = i
				break
			}
		}
		if selectedIndex == -1 {
			return fmt.Errorf("unknown soul selected: %s", option.Value)
		}

		detail := fmt.Sprintf("Soul: `%s`\nScope: %s\nCreated: %s\n\nContent:\n%s",
			selectedSoul.Name,
			strings.Join(selectedSoul.Scope, ", "),
			selectedSoul.CreatedAt.Format("2006-01-02 15:04:05"),
			selectedSoul.Content,
		)

		se := []databaseSoulPageAction{databaseSoulPageEdit, databaseSoulPageDelete, databaseSoulPageQuit}
		opts := cli.BuildOptions[databaseSoulPageAction](se, nil)
		form := newConfigForm[databaseSoulPageAction]()

		form.Register(databaseSoulPageEdit, func() error {
			contentRaw, err := p.Input("Soul Content", "Edit the system prompt content.", selectedSoul.Content, false, func(input string) error {
				if strings.TrimSpace(input) == "" {
					return fmt.Errorf("content is required")
				}
				return nil
			})
			if err != nil {
				return err
			}
			if err := repo.UpdateSoulContent(ctx, selectedSoul.Name, fmt.Sprintf("%v", contentRaw)); err != nil {
				return err
			}
			return p.Alert("Manboster Configuration Wizard", "Soul updated successfully!")
		})

		form.Register(databaseSoulPageDelete, func() error {
			confirm, err := p.Prompt(fmt.Sprintf("Do you want to delete soul %q? YOUR ACTION IS IRREVERSIBLE!", selectedSoul.Name), "Do you want to continue?", "Yes", "No")
			if err != nil {
				return err
			}
			if !confirm {
				return errQuit
			}
			if err := repo.DeleteSoul(ctx, selectedSoul.Name); err != nil {
				return err
			}
			if err := p.Alert("Manboster Configuration Wizard", "Soul deleted successfully!"); err != nil {
				return err
			}
			return errQuit
		})

		form.Register(databaseSoulPageQuit, nilFunc)

		err = handleWithPrompt[databaseSoulPageAction](p, form, opts, detail, "What do you want to do with this soul?")
		if err != nil {
			return err
		}
	}
}

func runDatabaseSoulCreate(ctx context.Context, p cli.Provider, repo repository.Repository) error {
	nameRaw, err := p.Input("Soul Name", "Enter a unique name for this soul.", "", false, func(input string) error {
		if strings.TrimSpace(input) == "" {
			return fmt.Errorf("name is required")
		}
		return nil
	})
	if err != nil {
		return err
	}
	name := fmt.Sprintf("%v", nameRaw)

	contentRaw, err := p.Input("Soul Content", "Enter the system prompt content for this soul.", "", false, func(input string) error {
		if strings.TrimSpace(input) == "" {
			return fmt.Errorf("content is required")
		}
		return nil
	})
	if err != nil {
		return err
	}
	content := fmt.Sprintf("%v", contentRaw)

	scopeRaw, err := p.Input("Scope", "Enter comma-separated scopes (e.g. global,telegram).", "global", false, func(input string) error {
		return nil
	})
	if err != nil {
		return err
	}
	var scope []string
	for _, s := range strings.Split(fmt.Sprintf("%v", scopeRaw), ",") {
		s = strings.TrimSpace(s)
		if s != "" {
			scope = append(scope, s)
		}
	}

	if err := repo.CreateSoul(ctx, types.Soul{
		Name:    name,
		Content: content,
		Scope:   scope,
	}); err != nil {
		return err
	}
	return p.Alert("Manboster Configuration Wizard", "Soul created successfully!")
}
