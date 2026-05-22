package file

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/engine/hook"
)

func (s *Service) RegisterHook(registry *hook.Registry) {
	registry.Register(hook.EngineAfterCompact, metadata.Name, hook.EngineAfterCompactHookProvider{
		PolyfillFunc: func(ctx context.Context, before string, after string) error {
			oldPath := config.Path(filepath.Join("workspace", "session-"+before))
			newPath := config.Path(filepath.Join("workspace", "session-"+after))

			if _, err := os.Stat(oldPath); os.IsNotExist(err) {
				return nil
			} else if err != nil {
				return fmt.Errorf("failed to stat old session dir %s: %w", oldPath, err)
			}

			if _, err := os.Stat(newPath); err == nil {
				if err := os.RemoveAll(newPath); err != nil {
					return fmt.Errorf("failed to remove existing new session dir %s: %w", newPath, err)
				}
			} else if !os.IsNotExist(err) {
				return fmt.Errorf("failed to stat new session dir %s: %w", newPath, err)
			}

			if err := os.Rename(oldPath, newPath); err != nil {
				return fmt.Errorf("failed to move session dir from %s to %s: %w", oldPath, newPath, err)
			}

			return nil
		},
	})
}
