package browser

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
)

var runWebpageInfo = tool.FactoryRegisterInfo[NameType]{
	Meta: schema.MetaData{
		Name:         "webpage",
		DisplayName:  "Fetch Webpage",
		Description:  "Fetch and return the content of a webpage by URL.",
		Represent:    "🌐",
		Irreversible: false,
	},
	Args:           schema.ArgsFromStruct(WebpageArgs{}),
	Run:            runWebpage,
	Continue:       tool.NilContinueFunc,
	CacheGroup:     tool.NilCacheGroupFunc,
	ClientRenderer: clientRendererWebpage,
}

func runWebpage(ctx context.Context, args string) (*plugin.RunResponse, error) {
	arg, err := util.Unmarshal[WebpageArgs](args)
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
	res, err := svc.ScrapWebpage(ctx, arg.URL, arg.ScrapType, arg.ResponseType, sessID)
	if err != nil {
		return nil, err
	}
	return &plugin.RunResponse{Response: res}, nil
}

func clientRendererWebpage(args string) string {
	arg, err := util.Unmarshal[WebpageArgs](args)
	if err != nil {
		return ""
	}
	return arg.URL
}
