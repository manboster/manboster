package datetime

import (
	"context"
	"time"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runTimeInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "time",
		DisplayName:  i18n.T(keys.TimeGetDisplayName),
		Description:  i18n.T(keys.TimeGetDescription),
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
