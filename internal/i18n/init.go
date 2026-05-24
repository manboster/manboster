package i18n

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
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
	// walk for subdirectories
	err := fs.WalkDir(localeFS, "locales", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// process json file
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".json") {
			_, loadErr := bundle.LoadMessageFileFS(localeFS, path)
			if loadErr != nil {
				color.Yellow(fmt.Sprintf("[Manboster i18n] Warning: failed to load translation file %s: %v", path, loadErr))
			}
		}
		return nil
	})

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
