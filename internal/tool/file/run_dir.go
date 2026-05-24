package file

import (
	"context"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runDirInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "dir",
		DisplayName:  i18n.T(keys.FileDirDisplayName),
		Description:  i18n.T(keys.FileDirDescription),
		Represent:    "🗂️",
		Irreversible: false,
	},
	Args:           schema.ArgsFromStruct(DirArgs{}),
	Run:            runDir,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: clientRendererDirPath,
}

func runDir(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, err := util.Unmarshal[DirArgs](args)
	if err != nil {
		return nil, err
	}
	pwd, err := resolvePwd(ctx, arg.IsPublic)
	if err != nil {
		return nil, err
	}
	return &plugin.RunResponse{Response: pwd}, nil
}
