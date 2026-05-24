package file

import (
	"context"
	"fmt"
	"os"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runDeleteInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "delete",
		DisplayName:  i18n.T(keys.FileDeleteDisplayName),
		Description:  i18n.T(keys.FileDeleteDescription),
		Represent:    "🗑️",
		Irreversible: true,
	},
	Args:           schema.ArgsFromStruct(DeleteArgs{}),
	Run:            runDelete,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: clientRendererFileName,
}

func runDelete(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, err := util.Unmarshal[DeleteArgs](args)
	if err != nil {
		return nil, err
	}
	if cfg.Mode == "readonly" {
		return nil, fmt.Errorf("failed to delete: read-only mode set by user")
	}
	pwd, err := resolvePwd(ctx, arg.IsPublic)
	if err != nil {
		return nil, err
	}
	sPath, err := getSafePath(pwd, arg.FilePath, arg.FileName)
	if err != nil {
		return nil, err
	}
	if err := os.Remove(sPath); err != nil {
		return nil, fmt.Errorf("failed to delete file %s: %w", sPath, err)
	}
	return &plugin.RunResponse{Response: "success"}, nil
}
