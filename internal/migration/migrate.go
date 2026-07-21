package migration

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/cli/helper"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/migration/patches"
	_ "github.com/manboster/manboster/internal/migration/patches/all"
	"github.com/spf13/cobra"
)

func MigrateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Migrate older application data",
		Run:   migrateCmd,
	}
}

func migrateCmd(cmd *cobra.Command, args []string) {
	color.Blue(i18n.T(keys.MigrationWizard))
	color.Blue(i18n.T(keys.MigrationStarted))

	err := config.Init()
	if err != nil {
		color.Yellow("[Manboster Config] No Configuration detected")
	}
	cfg := config.Read()

	p := patches.Patches()
	var patchesList []patches.Patch
	for _, patch := range p {
		if patch.Detect(cfg) {
			patchesList = append(patchesList, patch)
		}
	}

	if len(patchesList) == 0 {
		color.Green(i18n.T(keys.MigrationNone))
		return
	}

	color.Cyan(i18n.T(keys.MigrationDetected, len(patchesList)))
	for _, patch := range patchesList {
		color.Yellow(fmt.Sprintf("%s, introduced in %s", i18n.T(patch.Description), patch.Introduce))
	}

	color.Blue(i18n.T(keys.MigrationConfirm))
	chr := helper.GetChar()
	if strings.ToLower(chr) == "y" {
		newconf := cfg
		for _, patch := range patchesList {
			newconf, err = patch.Migrate(cfg)
			if err != nil {
				color.Green(i18n.Te(keys.MigrationFailed, "", err))
				continue
			}
		}

		err := config.Write(newconf, config.Path(""))
		if err != nil {
			color.Green(i18n.Te(keys.MigrationFailed, "", err))
			return
		}

		color.Red(i18n.T(keys.MigrationSuccess))
	}

	color.Yellow(i18n.Te(keys.MigrationFailed, "", errors.New("user canceled")))
	os.Exit(1)
}
