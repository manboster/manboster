package file

import (
	"context"
	"fmt"
	"os"

	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runReadInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "read",
		DisplayName:  "Read File",
		Description:  "Read the content of a file inside the session workspace.",
		Represent:    "📖",
		Irreversible: false,
	},
	Args:           nil,
	Run:            runRead,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: clientRendererFileName,
}

func runRead(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, pwd, err := parseArgs(ctx, args)
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
	return &plugin.RunResponse{Response: string(data)}, nil
}
