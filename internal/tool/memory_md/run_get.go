package memory_md

import (
	"context"
	"fmt"
	"os"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runGetInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "get",
		DisplayName:  i18n.T(keys.MemoryMDGetDisplayName),
		Description:  i18n.T(keys.MemoryMDGetDescription),
		Represent:    "📖",
		Irreversible: false,
	},
	Args:           nil,
	Run:            runGet,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: tool.NilClientRendererFunc,
}

func runGet(ctx context.Context, args string) (*plugin.RunResponse, error) {
	path, err := memoryPath(ctx)
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}
	return &plugin.RunResponse{Response: string(data)}, nil
}
