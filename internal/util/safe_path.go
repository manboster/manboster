package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func SafePath(baseDir, targetPath string) (string, error) {
	baseDir, err := filepath.EvalSymlinks(baseDir)
	if err != nil {
		return "", fmt.Errorf("failed to get basepath: %w", err)
	}

	absPath, err := filepath.Abs(filepath.Join(baseDir, targetPath))
	if err != nil {
		return "", fmt.Errorf("failed to parse path: %w", err)
	}

	realPath, err := filepath.EvalSymlinks(absPath)
	if err != nil {
		if os.IsNotExist(err) {
			cleanPath := filepath.Clean(absPath)
			rel, relErr := filepath.Rel(baseDir, cleanPath)
			if relErr != nil || strings.HasPrefix(rel, "..") {
				return "", fmt.Errorf("path overflow")
			}
			return absPath, nil
		}
		return "", fmt.Errorf("failed to parse path: %w", err)
	}

	rel, err := filepath.Rel(baseDir, realPath)
	if err != nil || strings.HasPrefix(rel, "..") {
		return "", fmt.Errorf("path overflow")
	}

	return absPath, nil
}
