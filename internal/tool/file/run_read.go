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

var runReadInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "read",
		DisplayName:  i18n.T(keys.FileReadDisplayName),
		Description:  i18n.T(keys.FileReadDescription),
		Represent:    "📖",
		Irreversible: false,
	},
	Args:           schema.ArgsFromStruct(ReadArgs{}),
	Run:            runRead,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: clientRendererFileName,
}

func runRead(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, err := util.Unmarshal[ReadArgs](args)
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
	data, err := os.ReadFile(sPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", sPath, err)
	}
	return &plugin.RunResponse{Response: string(data)}, nil
}
