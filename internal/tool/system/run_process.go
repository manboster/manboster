package system

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runProcessListInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "process-list",
		DisplayName:  "List Processes",
		Description:  "List all running processes sorted by CPU usage.",
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
		DisplayName:  "Process Info",
		Description:  "Get detailed information about a specific process by PID.",
		Represent:    "🔎",
		Irreversible: false,
	},
	Args:           schema.ArgsFromStruct(ProcessInfoArgs{}),
	Run:            runProcessInfo,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: tool.NilClientRendererFunc,
}

var runProcessKillInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "process_kill",
		DisplayName:  "Kill Process",
		Description:  "Kill a running process by PID.",
		Represent:    "🔴",
		Irreversible: true,
	},
	Args:           schema.ArgsFromStruct(ProcessKillArgs{}),
	Run:            runProcessKill,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: tool.NilClientRendererFunc,
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
