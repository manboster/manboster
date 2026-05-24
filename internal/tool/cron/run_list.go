package cron

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runListInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "list",
		DisplayName:  i18n.T(keys.CronListDisplayName),
		Description:  i18n.T(keys.CronListDescription),
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
	chatID, ok := ctx.Value("chat_id").(string)
	if !ok {
		return nil, fmt.Errorf("chat_id not found in context")
	}
	chatProvider, ok := ctx.Value("chat_provider").(string)
	if !ok {
		return nil, fmt.Errorf("chat_provider not found in context")
	}
	list, err := svc.List(ctx, chatProvider, chatID)
	if err != nil {
		return nil, err
	}
	jsonify, err := json.Marshal(list)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal list: %w", err)
	}
	return &plugin.RunResponse{Response: string(jsonify)}, nil
}
