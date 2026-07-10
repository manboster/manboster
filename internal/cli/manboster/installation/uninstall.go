package installation

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/cli/helper"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/spf13/cobra"
)

func UninstallCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstall all Manboster components with data",
		Run:   uninstallCmd,
	}
}

func uninstallCmd(cmd *cobra.Command, args []string) {
	color.Red(i18n.T(keys.AppUninstallPrompt))
	chr := helper.GetChar()
	if strings.ToLower(chr) == "y" {
		switch config.PackageType(config.Package) {
		case config.PackageAUR:
			color.Yellow(i18n.Te(keys.AppUninstallError, "", errors.New("you installed Manboster via AUR, use `paru -R manboster-bin` or `yay -R manboster-bin` to uninstall")))
			os.Exit(1)
		case config.PackageAOSC:
			color.Yellow(i18n.Te(keys.AppUninstallError, "", errors.New("you installed Manboster via oma, use `oma remove manboster` to uninstall")))
			os.Exit(1)
		case config.PackageBrew:
			color.Yellow(i18n.Te(keys.AppUninstallError, "", errors.New("you installed Manboster via brew, use `brew remove manboster` to uninstall")))
			os.Exit(1)
		default:
		}

		color.Red(i18n.T(keys.AppUninstallDataPrompt))
		chrData := helper.GetChar()
		if strings.ToLower(chrData) == "y" {
			err := reset()
			if err != nil {
				color.Green(i18n.Te(keys.AppResetError, "", err))
				os.Exit(1)
			}
		}

		executable, err := os.Executable()
		if err != nil {
			color.Red(i18n.T(keys.AppUninstallError, "", err))
			os.Exit(1)
			return
		}

		if runtime.GOOS == "windows" {
			cmdStr := fmt.Sprintf("ping 127.0.0.1 -n 2 > nul & del /f /q \"%s\"", executable)
			cmd := exec.Command("cmd", "/c", cmdStr)

			err := cmd.Start()
			if err != nil {
				color.Red(i18n.T(keys.AppUninstallError, "", err))
				os.Exit(1)
			}

			os.Exit(0)
		}

		err = os.Remove(executable)
		if err != nil {
			color.Red(i18n.T(keys.AppUninstallError, "", err))
			os.Exit(1)
		}

		color.Red(i18n.T(keys.AppUninstallSuccess))
		os.Exit(0)
	}

	color.Yellow(i18n.Te(keys.AppUninstallError, "", errors.New("user canceled")))
	os.Exit(1)
}
