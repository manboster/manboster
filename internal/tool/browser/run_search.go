package browser

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

var runSearchInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "search",
		DisplayName:  i18n.T(keys.BrowserSearchDisplayName),
		Description:  i18n.T(keys.BrowserSearchDescription),
		Represent:    "🔍",
		Irreversible: false,
	},
	Args:           schema.ArgsFromStruct(SearchArgs{}),
	Run:            runSearch,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: clientRendererSearch,
}

func runSearch(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, err := util.Unmarshal[SearchArgs](args)
	if err != nil {
		return nil, err
	}
	if !svc.isReady {
		return nil, fmt.Errorf("the browser is not ready, please wait for a while or check out whether there is an error or not")
	}
	sessID, ok := ctx.Value("session_id").(string)
	if !ok {
		return nil, fmt.Errorf("session_id not found in context")
	}
	res, err := svc.doWebSearch(ctx, arg.Keywords, arg.Engine, ResponseTypeRaw, sessID)
	if err != nil {
		return nil, err
	}
	return &plugin.RunResponse{Response: res}, nil
}

func clientRendererSearch(args string) string {
	arg, err := util.Unmarshal[SearchArgs](args)
	if err != nil {
		return ""
	}
	return arg.Keywords + " using " + string(arg.Engine)
}
