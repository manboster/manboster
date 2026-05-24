package file

import (
	"context"

	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runDirInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "dir",
		DisplayName:  "Get Workspace Dir",
		Description:  "Get the absolute path of the current session workspace directory.",
		Represent:    "🗂️",
		Irreversible: false,
	},
	Args:           nil,
	Run:            runDir,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: clientRendererDirPath,
}

func runDir(ctx context.Context, args string) (*plugin.RunResponse, error) {
	_, pwd, err := parseArgs(ctx, args)
	if err != nil {
		return nil, err
	}
	return &plugin.RunResponse{Response: pwd}, nil
}
