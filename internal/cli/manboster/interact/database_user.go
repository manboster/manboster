package interact

import (
	"context"
	"fmt"

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
		return "Promote to Admin"
	case databaseUserPageDegrade:
		return "Degrade to Unknown"
	case databaseUserPageDelete:
		return "Delete this User"
	case databaseUserPageQuit:
		return "Quit"
	default:
		return ""
	}
}

func runDatabaseUserConfig(p cli.Provider, repo repository.Repository) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	users, err := repo.GetAllUsers(ctx)
	if err != nil {
		return err
	}

	// list all users, then quit
	options := []cli.Option{}
	for _, user := range users {
		options = append(options, cli.Option{
			Key:   fmt.Sprintf("`%s` (%s) type: %s, created at %s", user.UserID, user.Platform, user.Type.String(), user.CreatedAt.Format("2006-01-02 15:04:05")),
			Value: fmt.Sprintf("%s:%s", user.Platform, user.UserID),
		})
	}
	options = append(options, quitOption)

	summary := fmt.Sprintf("%d users registered.", len(users))

	var option cli.Option
	for {
		option, err = p.Select("Select a user to manage.", summary, options, option.Value, func(o cli.Option) error {
			return nil
		})
		if err != nil {
			return err
		}

		if option.Value == _QUIT_ {
			return nil
		}

		if len(users) == 0 {
			if err := p.Alert("Manboster Configuration Wizard", "No users found!"); err != nil {
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
			confirm, err := p.Prompt(fmt.Sprintf("Promote user %q to Admin?", selectedUser.UserID), "Do you want to continue?", "Yes", "No")
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
			return p.Alert("Manboster Configuration Wizard", "User promoted to Admin!")
		})

		form.Register(databaseUserPageDegrade, func() error {
			confirm, err := p.Prompt(fmt.Sprintf("Degrade user %q to Unknown?", selectedUser.UserID), "Do you want to continue?", "Yes", "No")
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
			return p.Alert("Manboster Configuration Wizard", "User degraded to Unknown!")
		})

		form.Register(databaseUserPageDelete, func() error {
			confirm, err := p.Prompt(fmt.Sprintf("Do you want to delete user %q? YOUR ACTION IS IRREVERSIBLE!", selectedUser.UserID), "Do you want to continue?", "Yes", "No")
			if err != nil {
				return err
			}
			if !confirm {
				return errQuit
			}
			if err := repo.DeleteUser(ctx, selectedUser.Platform, selectedUser.UserID); err != nil {
				return err
			}
			return p.Alert("Manboster Configuration Wizard", "User deleted successfully!")
		})

		form.Register(databaseUserPageQuit, nilFunc)

		err = handleWithPrompt[databaseUserPageAction](p, form, opts, detail, "What do you want to do with this user?")
		if err != nil {
			return err
		}
	}
}
