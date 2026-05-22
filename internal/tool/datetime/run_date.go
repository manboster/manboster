package datetime

import (
	"context"
	"time"

	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runDateInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "date",
		DisplayName:  "Get date",
		Description:  "Get this machine's current date in format like 2026-01-11",
		Represent:    "📅",
		Irreversible: false,
	},
	Args:           nil,
	Run:            runDate,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: tool.NilClientRendererFunc,
}

func runDate(ctx context.Context, args string) (*plugin.RunResponse, error) {
	resp := &plugin.RunResponse{
		Hangup: false,
	}
	resp.Response = time.Now().Format("2006-01-02")
	return resp, nil
}
