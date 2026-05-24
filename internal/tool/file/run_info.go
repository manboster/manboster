package file

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runInfoInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "info",
		DisplayName:  "File Info",
		Description:  "Get metadata of a file or directory inside the session workspace.",
		Represent:    "ℹ️",
		Irreversible: false,
	},
	Args:           schema.ArgsFromStruct(InfoArgs{}),
	Run:            runInfo,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: clientRendererFileName,
}

func runInfo(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, err := util.Unmarshal[InfoArgs](args)
	if err != nil {
		return nil, err
	}
	pwd, err := resolvePwd(ctx, arg.IsPublic)
	if err != nil {
		return nil, err
	}
	sPath, err := getSafePath(pwd, arg.FilePath, arg.FileName)
	if err != nil {
		return nil, err
	}
	info, err := os.Stat(sPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get file info %s: %w", sPath, err)
	}
	jsonify, err := json.Marshal(newFileEntry(info))
	if err != nil {
		return nil, fmt.Errorf("failed to jsonify file info %s: %w", sPath, err)
	}
	return &plugin.RunResponse{Response: string(jsonify)}, nil
}
