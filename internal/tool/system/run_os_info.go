package system

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

var runOSInfoInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "os-info",
		DisplayName:  i18n.T(keys.SystemOSInfoDisplayName),
		Description:  i18n.T(keys.SystemOSInfoDescription),
		Represent:    "🖥️",
		Irreversible: false,
	},
	Args:           nil,
	Run:            runOSInfo,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: tool.NilClientRendererFunc,
}

func runOSInfo(ctx context.Context, args string) (*plugin.RunResponse, error) {
	sys, err := getSystemInfo(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get system info: %w", err)
	}
	jsonify, err := json.Marshal(sys)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal system info: %w", err)
	}
	return &plugin.RunResponse{Response: string(jsonify)}, nil
}
