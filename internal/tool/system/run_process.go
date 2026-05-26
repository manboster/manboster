package system

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runProcessListInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "process-list",
		DisplayName:  i18n.T(keys.SystemProcessListDisplayName),
		Description:  i18n.T(keys.SystemProcessListDescription),
		Represent:    "📋",
		Irreversible: false,
	},
	Args:           nil,
	Run:            runProcessList,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: tool.NilClientRendererFunc,
}

var runProcessInfoInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "process-info",
		DisplayName:  i18n.T(keys.SystemProcessInfoDisplayName),
		Description:  i18n.T(keys.SystemProcessInfoDescription),
		Represent:    "🔎",
		Irreversible: false,
	},
	Args:           schema.ArgsFromStruct(ProcessInfoArgs{}),
	Run:            runProcessInfo,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: renderProcessInfo,
}

var runProcessKillInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "process-kill",
		DisplayName:  i18n.T(keys.SystemProcessKillDisplayName),
		Description:  i18n.T(keys.SystemProcessKillDescription),
		Represent:    "🔴",
		Irreversible: true,
	},
	Args:           schema.ArgsFromStruct(ProcessKillArgs{}),
	Run:            runProcessKill,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: renderProcessKill,
}

func runProcessList(ctx context.Context, args string) (*plugin.RunResponse, error) {
	list, err := listProcesses(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get process list: %w", err)
	}
	jsonify, err := json.Marshal(list)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal process list: %w", err)
	}
	return &plugin.RunResponse{Response: string(jsonify)}, nil
}

func runProcessInfo(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, err := util.Unmarshal[ProcessInfoArgs](args)
	if err != nil {
		return nil, err
	}
	info, err := getProcessInfo(ctx, int32(arg.PID))
	if err != nil {
		return nil, fmt.Errorf("failed to get process info: %w", err)
	}
	jsonify, err := json.Marshal(info)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal process info: %w", err)
	}
	return &plugin.RunResponse{Response: string(jsonify)}, nil
}

func runProcessKill(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, err := util.Unmarshal[ProcessKillArgs](args)
	if err != nil {
		return nil, err
	}
	if err := killProcess(ctx, int32(arg.PID)); err != nil {
		return nil, fmt.Errorf("failed to kill process: %w", err)
	}
	return &plugin.RunResponse{Response: "success"}, nil
}

func renderProcessInfo(args string) string {
	arg, err := util.Unmarshal[ProcessInfoArgs](args)
	if err != nil {
		return ""
	}
	return strconv.Itoa(arg.PID)
}

func renderProcessKill(args string) string {
	arg, err := util.Unmarshal[ProcessKillArgs](args)
	if err != nil {
		return ""
	}
	return strconv.Itoa(arg.PID)
}
