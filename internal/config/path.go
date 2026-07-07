package config

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
)

// Path gets the correct path of manboster file storage
func Path(filename string) string {
	_ = godotenv.Load()
	dir := os.Getenv("MANBOSTER_HOME")

	var p string
	if dir == "" {
		var err error
		p, err = os.UserConfigDir()
		if err != nil {
			return filename
		}
	} else {
		p = dir
	}

	// if there is none, create one.
	if _, err := os.ReadDir(filepath.Join(p, "manboster")); err != nil {
		err = os.MkdirAll(filepath.Join(p, "manboster"), 0700)
		if err != nil {
			return filename
		}
	}

	p = filepath.Join(p, "manboster", filename)

	return p
}

func LookupOlderPaths() {
	_ = godotenv.Load()
	dir := os.Getenv("MANBOSTER_HOME")
	p := ""
	if dir == "" {
		p, _ = os.UserHomeDir()
	} else {
		p = dir
	}

	if _, err := os.ReadDir(filepath.Join(p, ".manboster")); err == nil {
		color.Yellow(i18n.T(keys.AppDataWarning))
		color.Yellow(i18n.T(keys.AppInputPrompt))
		_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
