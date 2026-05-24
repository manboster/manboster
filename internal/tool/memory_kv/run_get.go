package memory_kv

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runGetInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "get",
		DisplayName:  "Get Memory",
		Description:  "Retrieve a value from memory by key.",
		Represent:    "🔑",
		Irreversible: false,
	},
	Args:           schema.ArgsFromStruct(GetArgs{}),
	Run:            runGet,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: tool.NilClientRendererFunc,
}

func runGet(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, err := util.Unmarshal[GetArgs](args)
	if err != nil {
		return nil, err
	}
	memory, err := memDB.GetMemory(ctx, arg.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to get %q", arg.Key)
	}
	return &plugin.RunResponse{Response: memory.Value}, nil
}
