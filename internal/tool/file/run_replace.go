package file

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runReplaceInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "replace",
		DisplayName:  "Replace in File",
		Description:  "Replace occurrences of a specific text in a file. Set line to 0 to replace all occurrences, or specify a line number to restrict replacement to that line.",
		Represent:    "🔄",
		Irreversible: true,
	},
	Args:           schema.ArgsFromStruct(ReplaceArgs{}),
	Run:            runReplace,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: clientRendererFileName,
}

func runReplace(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, err := util.Unmarshal[ReplaceArgs](args)
	if err != nil {
		return nil, err
	}
	if cfg.Mode == "readonly" {
		return nil, fmt.Errorf("failed to replace: read-only mode set by user")
	}
	pwd, err := resolvePwd(ctx, arg.IsPublic)
	if err != nil {
		return nil, err
	}
	sPath, err := getSafePath(pwd, arg.FilePath, arg.FileName)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(sPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", sPath, err)
	}

	var updated string
	if arg.Line == 0 {
		updated = strings.ReplaceAll(string(data), arg.OldText, arg.NewText)
		if updated == string(data) {
			return nil, fmt.Errorf("old_text not found in file %s", sPath)
		}
	} else {
		var sb strings.Builder
		scanner := bufio.NewScanner(strings.NewReader(string(data)))
		lineNum := 0
		replaced := false
		for scanner.Scan() {
			lineNum++
			line := scanner.Text()
			if lineNum == arg.Line {
				newLine := strings.ReplaceAll(line, arg.OldText, arg.NewText)
				if newLine == line {
					return nil, fmt.Errorf("old_text not found on line %d of file %s", arg.Line, sPath)
				}
				sb.WriteString(newLine)
				replaced = true
			} else {
				sb.WriteString(line)
			}
			sb.WriteByte('\n')
		}
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("failed to scan file %s: %w", sPath, err)
		}
		if !replaced {
			return nil, fmt.Errorf("line %d not found in file %s", arg.Line, sPath)
		}
		updated = sb.String()
	}

	if err := os.WriteFile(sPath, []byte(updated), 0644); err != nil {
		return nil, fmt.Errorf("failed to write file %s: %w", sPath, err)
	}
	return &plugin.RunResponse{Response: "success"}, nil
}
