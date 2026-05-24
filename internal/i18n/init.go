package i18n

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/jeandeaual/go-locale"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func Init() error {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// load embedded locales
	entries, err := fs.ReadDir(localeFS, "locales")
	if err != nil {
		return err
	}

	// go for locales files
	for _, entry := range entries {
		// check out json files
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {
			filePath := filepath.Join("locales", entry.Name())

			_, err := bundle.LoadMessageFileFS(localeFS, filePath)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster i18n] Error loading i18n: %v", err))
				continue
			}
		}
	}

	userLocales, err := locale.GetLocales()
	if err != nil {
		userLocales = []string{"en"}
	}

	if envLang := os.Getenv("MANBOSTER_LANG"); envLang != "" {
		userLocales = append([]string{envLang}, userLocales...)
	}

	localizer = i18n.NewLocalizer(bundle, userLocales...)
	return nil
}
