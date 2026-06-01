package loader

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/skill"
)

// LoadSkills loads all skills in path recursively
func LoadSkills(path string) error {
	truePath := path
	entries, err := os.ReadDir(truePath)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Loader] Could not read directory: %q", err))
		return err
	}

	fileMap := make(map[string]bool)

	for _, entry := range entries {
		identifyName := entry.Name()
		if entry.IsDir() {
			identifyName += ".md"
		}

		if fileMap[identifyName] {
			color.Yellow(fmt.Sprintf("[Manboster Loader] Duplicated skill file and folders! Skipped %q!", entry.Name()))
			continue
		}

		fileMap[identifyName] = true
		if entry.IsDir() {
			err = skill.Load(filepath.Join(truePath, entry.Name()), entry.Name(), true)
			if err != nil {
				color.Red(fmt.Sprintf("[Manboster Loader] Could not load skill %q: %q", entry.Name(), err))
				continue
			}
		} else {
			if strings.HasSuffix(entry.Name(), ".md") {
				err = skill.Load(filepath.Join(truePath, entry.Name()), strings.TrimSuffix(entry.Name(), ".md"), false)
				if err != nil {
					color.Red(fmt.Sprintf("[Manboster Loader] Could not load skill %q: %q", entry.Name(), err))
					continue
				}
			} else {
				color.Yellow(fmt.Sprintf("[Manboster Loader] Skipped %q because it's not a valid markdown file!", entry.Name()))
				continue
			}
		}
	}
	return nil
}
