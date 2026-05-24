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
	Args:           nil,
	Run:            runWrite,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: clientRendererFileName,
}

func runWrite(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, pwd, err := parseArgs(ctx, args)
	if err != nil {
		return nil, err
	}

	if cfg.Mode == "readonly" {
		return nil, fmt.Errorf("failed to write: read-only mode set by user")
	}

	sPath, err := getSafePath(pwd, arg.FilePath, arg.FileName)
	if err != nil {
		return nil, err
	}
	if err := os.WriteFile(sPath, []byte(arg.Content), 0644); err != nil {
		color.Yellow(fmt.Sprintf("[Manboster Tool Provider] dev.manboster.file failed to write file: %q", err))
		return nil, fmt.Errorf("failed to write file %s: %w", sPath, err)
	}
	return &plugin.RunResponse{Response: "Success"}, nil
}
