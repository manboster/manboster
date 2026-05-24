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
	Args:           schema.ArgsFromStruct(ListArgs{}),
	Run:            runList,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: clientRendererDirPath,
}

func runList(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, err := unmarshal[ListArgs](args)
	if err != nil {
		return nil, err
	}
	pwd, err := resolvePwd(ctx, arg.IsPublic)
	if err != nil {
		return nil, err
	}
	sPath, err := getSafePath(pwd, arg.FilePath, "")
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
