package file

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runWriteInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "write",
		DisplayName:  "Write File",
		Description:  "Write content to a file inside the session workspace.",
		Represent:    "✏️",
		Irreversible: true,
	},
	Args:           schema.ArgsFromStruct(WriteArgs{}),
	Run:            runWrite,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: clientRendererFileName,
}

func runWrite(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, err := unmarshal[WriteArgs](args)
	if err != nil {
		return nil, err
	}
	if cfg.Mode == "readonly" {
		return nil, fmt.Errorf("failed to write: read-only mode set by user")
	}
	pwd, err := resolvePwd(ctx, arg.IsPublic)
	if err != nil {
		return nil, err
	}
	sPath, err := getSafePath(pwd, arg.FilePath, arg.FileName)
	if err != nil {
		return nil, err
	}

	if arg.Append {
		f, err := os.OpenFile(sPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to open file %s: %w", sPath, err)
		}
		defer f.Close()
		if _, err := f.WriteString(arg.Content); err != nil {
			return nil, fmt.Errorf("failed to append to file %s: %w", sPath, err)
		}
	} else {
		if err := os.WriteFile(sPath, []byte(arg.Content), 0644); err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Tool Provider] dev.manboster.file failed to write file: %q", err))
			return nil, fmt.Errorf("failed to write file %s: %w", sPath, err)
		}
	}
	return &plugin.RunResponse{Response: "success"}, nil
}
