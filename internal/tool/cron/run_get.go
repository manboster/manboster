package cron

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runGetInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "get",
		DisplayName:  i18n.T(keys.CronGetDisplayName),
		Description:  i18n.T(keys.CronGetDescription),
		Represent:    "🔎",
		Irreversible: false,
	},
	Args:           schema.ArgsFromStruct(GetArgs{}),
	Run:            runGet,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: clientRendererJobName,
}

func runGet(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, err := util.Unmarshal[GetArgs](args)
	if err != nil {
		return nil, err
	}
	data, err := svc.Get(ctx, arg.JobName)
	if err != nil {
		return nil, err
	}
	jsonify, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal job: %w", err)
	}
	return &plugin.RunResponse{Response: string(jsonify)}, nil
}
