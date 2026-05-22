package datetime

import (
	"context"
	"time"

	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runTimeInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "time",
		DisplayName:  "Get time",
		Description:  "Get this machine's current time in format like 12:34:56+08",
		Represent:    "🕙",
		Irreversible: false,
	},
	Args:           nil,
	Run:            runTime,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: tool.NilClientRendererFunc,
}

func runTime(ctx context.Context, args string) (*plugin.RunResponse, error) {
	resp := &plugin.RunResponse{
		Hangup: false,
	}
	resp.Response = time.Now().Format("15:04:05-07")
	return resp, nil
}
