package cron

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runSetInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "set",
		DisplayName:  i18n.T(keys.CronSetDisplayName),
		Description:  i18n.T(keys.CronSetDescription),
		Represent:    "⏰",
		Irreversible: false,
	},
	Args:           schema.ArgsFromStruct(SetArgs{}),
	Run:            runSet,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: clientRendererJobName,
}

func runSet(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, err := util.Unmarshal[SetArgs](args)
	if err != nil {
		return nil, err
	}
	if arg.MessageType != MessagePrompt && arg.MessageType != MessageText {
		return nil, fmt.Errorf("invalid argument '%s' in messageType", arg.MessageType)
	}

	chatID, ok := ctx.Value("chat_id").(string)
	if !ok {
		return nil, fmt.Errorf("chat_id not found in context")
	}
	chatProvider, ok := ctx.Value("chat_provider").(string)
	if !ok {
		return nil, fmt.Errorf("chat_provider not found in context")
	}
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, fmt.Errorf("user_id not found in context")
	}

	if err := svc.Create(ctx, arg, chatID, chatProvider, userID); err != nil {
		return nil, err
	}
	return &plugin.RunResponse{Response: "success"}, nil
}

func clientRendererJobName(args string) string {
	type jobNameArgs struct {
		JobName string `json:"job_name"`
	}
	arg, err := util.Unmarshal[jobNameArgs](args)
	if err != nil {
		return ""
	}
	return arg.JobName
}
