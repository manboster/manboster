package migration

import "github.com/spf13/cobra"

func MigrateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Migrate older application data",
		Run:   migrateCmd,
	}
}

func migrateCmd(cmd *cobra.Command, args []string) {

}
