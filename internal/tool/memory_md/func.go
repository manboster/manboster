package memory_md

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/manboster/manboster/internal/config"
)

func memoryPath(ctx context.Context) (string, error) {
	chatID, ok := ctx.Value("chat_id").(string)
	if !ok {
		return "", fmt.Errorf("chat_id not found in context")
	}
	chatProvider, ok := ctx.Value("chat_provider").(string)
	if !ok {
		return "", fmt.Errorf("chat_provider not found in context")
	}
	return config.Path(filepath.Join("memory", fmt.Sprintf("memory-%s-%s.md", chatProvider, chatID))), nil
}
