package memory_md

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

type SetArgs struct {
	Value string `json:"value" description:"Markdown content to store." example:"# Notes\ncontent" validate:"required"`
}

var runSetInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "set",
		DisplayName:  i18n.T(keys.MemoryMDSetDisplayName),
		Description:  i18n.T(keys.MemoryMDSetDescription),
		Represent:    "✏️",
		Irreversible: false,
	},
	Args:           schema.ArgsFromStruct(SetArgs{}),
	Run:            runSet,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: tool.NilClientRendererFunc,
}

func runSet(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, err := util.Unmarshal[SetArgs](args)
	if err != nil {
		return nil, err
	}
	if len([]byte(arg.Value)) > 65536 {
		return nil, fmt.Errorf("failed to write: markdown entity too large")
	}
	path, err := memoryPath(ctx)
	if err != nil {
		return nil, err
	}
	if err := os.WriteFile(path, []byte(arg.Value), 0644); err != nil {
		return nil, fmt.Errorf("failed to write file %s: %w", path, err)
	}
	return &plugin.RunResponse{Response: "success"}, nil
}
