package interact

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/spec/cli"
)

type databaseUserPageAction string

const (
	databaseUserPagePromote databaseUserPageAction = "promote"
	databaseUserPageDegrade databaseUserPageAction = "degrade"
	databaseUserPageDelete  databaseUserPageAction = _DELETE_
	databaseUserPageQuit    databaseUserPageAction = _QUIT_
)

func (a databaseUserPageAction) Name() string { return string(a) }
func (a databaseUserPageAction) DisplayName() string {
	switch a {
	case databaseUserPagePromote:
		return i18n.T(keys.CliConfigUserPromoteAction)
	case databaseUserPageDegrade:
		return i18n.T(keys.CliConfigUserDegradeAction)
	case databaseUserPageDelete:
		return i18n.T(keys.CliConfigUserDeleteAction)
	case databaseUserPageQuit:
		return i18n.T(keys.CliConfigActionQuit)
	default:
		return ""
	}
}

func runDatabaseUserConfig(p cli.Provider, repo repository.Repository) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var option cli.Option
	for {
		// reload on every iteration so changes are reflected
		users, err := repo.GetAllUsers(ctx)
		if err != nil {
			return err
		}

		options := []cli.Option{}
		for _, user := range users {
			options = append(options, cli.Option{
				Key:   fmt.Sprintf("`%s` (%s) type: %s, created at %s", user.UserID, user.Platform, user.Type.String(), user.CreatedAt.Format("2006-01-02 15:04:05")),
				Value: fmt.Sprintf("%s:%s", user.Platform, user.UserID),
			})
		}
		options = append(options, quitOption)

		summary := fmt.Sprintf("%d users registered.", len(users))

		option, err = p.Select(i18n.T(keys.CliConfigUserSelectPrompt), summary, options, option.Value, func(o cli.Option) error {
			return nil
		})
		if err != nil {
			return err
		}

		if option.Value == _QUIT_ {
			return nil
		}

		if len(users) == 0 {
			if err := p.Alert(i18n.T(keys.CliWizardTitle), i18n.T(keys.CliConfigUserNoUsers)); err != nil {
				return err
			}
			return nil
		}

		// find selected user by "platform:userID"
		var selectedUser types.User
		selectedIndex := -1
		for i, user := range users {
			if fmt.Sprintf("%s:%s", user.Platform, user.UserID) == option.Value {
				selectedUser = user
				selectedIndex = i
				break
			}
		}
		if selectedIndex == -1 {
			return fmt.Errorf("unknown user selected: %s", option.Value)
		}

		detail := fmt.Sprintf("User: `%s`\nPlatform: %s\nType: %s\nCreated: %s\nUpdated: %s",
			selectedUser.UserID,
			selectedUser.Platform,
			selectedUser.Type.String(),
			selectedUser.CreatedAt.Format("2006-01-02 15:04:05"),
			selectedUser.UpdatedAt.Format("2006-01-02 15:04:05"),
		)

		se := []databaseUserPageAction{databaseUserPagePromote, databaseUserPageDegrade, databaseUserPageDelete, databaseUserPageQuit}
		opts := cli.BuildOptions[databaseUserPageAction](se, nil)
		form := newConfigForm[databaseUserPageAction]()

		form.Register(databaseUserPagePromote, func() error {
			confirm, err := p.Prompt(fmt.Sprintf(i18n.T(keys.CliConfigUserPromoteConfirm), selectedUser.UserID), "Do you want to continue?", "Yes", "No")
			if err != nil {
				return err
			}
			if !confirm {
				return errQuit
			}
			if err := repo.DeleteUser(ctx, selectedUser.Platform, selectedUser.UserID); err != nil {
				return err
			}
			if err := repo.CreateUser(ctx, types.User{
				UserID:   selectedUser.UserID,
				Platform: selectedUser.Platform,
				Type:     types.UserAdmin,
			}); err != nil {
				return err
			}
			return p.Alert(i18n.T(keys.CliWizardTitle), i18n.T(keys.CliConfigUserPromoteSuccess))
		})

		form.Register(databaseUserPageDegrade, func() error {
			confirm, err := p.Prompt(fmt.Sprintf(i18n.T(keys.CliConfigUserDegradeConfirm), selectedUser.UserID), "Do you want to continue?", "Yes", "No")
			if err != nil {
				return err
			}
			if !confirm {
				return errQuit
			}
			if err := repo.DeleteUser(ctx, selectedUser.Platform, selectedUser.UserID); err != nil {
				return err
			}
			if err := repo.CreateUser(ctx, types.User{
				UserID:   selectedUser.UserID,
				Platform: selectedUser.Platform,
				Type:     types.UserUnknown,
			}); err != nil {
				return err
			}
			return p.Alert(i18n.T(keys.CliWizardTitle), i18n.T(keys.CliConfigUserDegradeSuccess))
		})

		form.Register(databaseUserPageDelete, func() error {
			confirm, err := p.Prompt(i18n.Te(keys.CliConfigUserDeleteConfirm, selectedUser.Platform+":"+selectedUser.UserID, nil), "Do you want to continue?", "Yes", "No")
			if err != nil {
				return err
			}
			if !confirm {
				return errQuit
			}
			if err := repo.DeleteUser(ctx, selectedUser.Platform, selectedUser.UserID); err != nil {
				return err
			}
			if err := p.Alert(i18n.T(keys.CliWizardTitle), i18n.T(keys.CliConfigUserDeleteSuccess)); err != nil {
				return err
			}
			return errQuit
		})

		form.Register(databaseUserPageQuit, nilFunc)

		err = handleWithPrompt[databaseUserPageAction](p, form, opts, detail, i18n.T(keys.CliConfigActionWhatToDo))
		if err != nil {
			return err
		}
	}
}
