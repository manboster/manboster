package memory_md

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runGetInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "get",
		DisplayName:  "Get Markdown Memory",
		Description:  "Read the chat-specific markdown memory file.",
		Represent:    "📖",
		Irreversible: false,
	},
	Args:           nil,
	Run:            runGet,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: tool.NilClientRendererFunc,
}

var runSetInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "set",
		DisplayName:  "Set Markdown Memory",
		Description:  "Write markdown content to the chat-specific memory file. Maximum 64KB.",
		Represent:    "✏️",
		Irreversible: false,
	},
	Args:           schema.ArgsFromStruct(SetArgs{}),
	Run:            runSet,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: tool.NilClientRendererFunc,
}

func memoryPath(ctx context.Context) (string, error) {
	chatID, ok := ctx.Value("chat_id").(string)
	if !ok {
		return "", fmt.Errorf("chat_id not found in context")
	}
	chatProvider, ok := ctx.Value("chat_provider").(string)
	if !ok {
		return "", fmt.Errorf("chat_provider not found in context")
	}
	return config.Path(filepath.Join("memory", fmt.Sprintf("memory-%s-%s.md", chatProvider, chatID))), nil
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

func runSet(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, err := util.Unmarshal[SetArgs](args)
	if err != nil {
		return nil, err
	}
	if len([]byte(arg.Value)) > 65536 {
		return nil, fmt.Errorf("failed to write: markdown entity too large")
	}
	path, err := memoryPath(ctx)
	if err != nil {
		return nil, err
	}
	if err := os.WriteFile(path, []byte(arg.Value), 0644); err != nil {
		return nil, fmt.Errorf("failed to write file %s: %w", path, err)
	}
	return &plugin.RunResponse{Response: "success"}, nil
}
