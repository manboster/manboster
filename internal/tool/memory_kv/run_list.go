package memory_kv

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
		DisplayName:  "List Memory Keys",
		Description:  "List all stored memory keys.",
		Represent:    "📋",
		Irreversible: false,
	},
	Args:           nil,
	Run:            runList,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: tool.NilClientRendererFunc,
}

func runList(ctx context.Context, args string) (*plugin.RunResponse, error) {
	keys, err := memDB.ListMemoryKeys(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list keys")
	}
	jsonify, err := json.Marshal(keys)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal keys")
	}
	return &plugin.RunResponse{Response: string(jsonify)}, nil
}
