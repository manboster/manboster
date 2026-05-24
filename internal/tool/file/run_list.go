package file

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runListInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "list",
		DisplayName:  "List Directory",
		Description:  "List files and directories inside a path in the session workspace.",
		Represent:    "📂",
		Irreversible: false,
	},
	Args:           nil,
	Run:            runList,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: clientRendererDirPath,
}

func runList(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, pwd, err := parseArgs(ctx, args)
	if err != nil {
		return nil, err
	}

	if arg.FileName != "" {
		return nil, fmt.Errorf("filename is not allowed to give while name is list")
	}

	sPath, err := getSafePath(pwd, arg.FilePath, arg.FileName)
	if err != nil {
		return nil, err
	}
	d, err := listDir(sPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir %s: %w", sPath, err)
	}
	jsonify, err := json.Marshal(d)
	if err != nil {
		return nil, fmt.Errorf("failed to jsonify dir %s: %w", sPath, err)
	}
	return &plugin.RunResponse{Response: string(jsonify)}, nil
}
