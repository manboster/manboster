package memory_kv

import (
	"context"
	"errors"
	"fmt"

	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runSetInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "set",
		DisplayName:  "Set Memory",
		Description:  "Store a value in memory by key. Creates the key if it does not exist.",
		Represent:    "💾",
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
	err = memDB.EditMemoryValue(ctx, arg.Key, arg.Value)
	if errors.Is(err, repository.ErrNotFound) {
		err = memDB.CreateMemory(ctx, types.Memory{
			Key:   arg.Key,
			Value: arg.Value,
		})
	}
	if err != nil {
		return nil, fmt.Errorf("failed to store %q", arg.Key)
	}
	return &plugin.RunResponse{Response: "success"}, nil
}
