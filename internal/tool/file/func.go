package file

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/util"
)

func getSafePath(pwd string, filePath []string, filename string) (string, error) {
	dir := ""
	for _, p := range filePath {
		dir = filepath.Join(dir, p)
	}
	dir = filepath.Join(dir, filename)
	sPath, err := util.SafePath(pwd, dir)
	if err != nil {
		return "", fmt.Errorf("parsing path error: %w", err)
	}
	return sPath, nil
}

type fileEntry struct {
	Name    string    `json:"name"`
	Size    int64     `json:"size"`
	Mode    string    `json:"mode"`
	ModTime time.Time `json:"mod_time"`
	IsDir   bool      `json:"is_dir"`
}

func newFileEntry(info os.FileInfo) fileEntry {
	return fileEntry{
		Name:    info.Name(),
		Size:    info.Size(),
		Mode:    info.Mode().String(),
		ModTime: info.ModTime(),
		IsDir:   info.IsDir(),
	}
}

func listDir(path string) ([]fileEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	result := make([]fileEntry, 0, len(entries))
	for _, e := range entries {
		info, err := e.Info()
		if err != nil {
			continue
		}
		result = append(result, fileEntry{
			Name:    info.Name(),
			Size:    info.Size(),
			Mode:    info.Mode().String(),
			ModTime: info.ModTime(),
			IsDir:   info.IsDir(),
		})
	}
	return result, nil
}

func unmarshalArgs(args string) (RunArgs, error) {
	var arg RunArgs
	if err := json.Unmarshal([]byte(args), &arg); err != nil {
		return RunArgs{}, fmt.Errorf("invalid arguments")
	}
	return arg, nil
}

func clientRendererFileName(args string) string {
	arg, err := unmarshalArgs(args)
	if err != nil {
		return ""
	}
	return filepath.Join(filepath.Join(arg.FilePath...), arg.FileName)
}

func clientRendererDirPath(args string) string {
	arg, err := unmarshalArgs(args)
	if err != nil {
		return ""
	}
	return filepath.Join(arg.FilePath...)
}

func parseArgs(ctx context.Context, args string) (RunArgs, string, error) {
	arg, err := unmarshalArgs(args)
	if err != nil {
		return RunArgs{}, "", err
	}

	sessID, ok := ctx.Value("session_id").(string)
	if !ok {
		return RunArgs{}, "", fmt.Errorf("session_id not found in context")
	}
	pwd := config.Path(filepath.Join("workspace", "session-"+sessID))
	if err := os.MkdirAll(pwd, 0755); err != nil {
		return RunArgs{}, "", fmt.Errorf("failed to create session dir: %w", err)
	}

	if arg.IsPublic {
		pwd = config.Path(filepath.Join("workspace", "public"))
	}
	return arg, pwd, nil
}
