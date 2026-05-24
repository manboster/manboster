package cron

import (
	"context"

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
		DisplayName:  i18n.T(keys.CronDeleteDisplayName),
		Description:  i18n.T(keys.CronDeleteDescription),
		Represent:    "🗑️",
		Irreversible: true,
	},
	Args:           schema.ArgsFromStruct(DeleteArgs{}),
	Run:            runDelete,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: clientRendererJobName,
}

func runDelete(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, err := util.Unmarshal[DeleteArgs](args)
	if err != nil {
		return nil, err
	}
	if err := svc.Delete(ctx, arg.JobName); err != nil {
		return nil, err
	}
	return &plugin.RunResponse{Response: "success"}, nil
}
