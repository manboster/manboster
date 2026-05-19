package interact

import (
	"github.com/spf13/cobra"
)

func ConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config-next [command]",
		Short: "[Experimental] Run configuration wizard for Manboster application",
		Args:  cobra.MaximumNArgs(1),
		Run:   configCmdRun,
	}

	cmd.Flags().BoolP("edit", "e", false, "Open config file in terminal $EDITOR")

	editCmd := &cobra.Command{
		Use:   "edit",
		Short: "Open config file in terminal $EDITOR",
		Args:  cobra.NoArgs,
		Run:   configCmdEditRun,
	}
	openCmd := &cobra.Command{
		Use:   "open",
		Short: "Open config file using your system's default editor",
		Args:  cobra.NoArgs,
		Run:   configCmdOpenRun,
	}
	helpCmd := &cobra.Command{
		Use:   "help",
		Short: "The help of this subcommand",
		Args:  cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			if c.HasParent() {
				err := c.Parent().Help()
				if err != nil {
					return
				}
			}
		},
	}

	cmd.AddCommand(helpCmd)
	cmd.AddCommand(editCmd)
	cmd.AddCommand(openCmd)

	return cmd
}

// OnboardConfigCmd provides an interactive TUI to configure your manboster application.
func OnboardConfigCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "onboard",
		Short: "Run onboard configuration wizard for Manboster application",
		Run:   OnboardConfigCmdRun,
	}
}
