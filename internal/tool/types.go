package tool

import (
	"context"

	"github.com/manboster/manboster/spec/plugin"
)

type RunFunc func(ctx context.Context, args string) (*plugin.RunResponse, error)
type ContinueFunc func(ctx context.Context, session string) (*plugin.RunResponse, error)
type CacheGroupFunc func(args string) string
type ClientRendererFunc func(args string) string

var NilRunFunc RunFunc = func(ctx context.Context, args string) (*plugin.RunResponse, error) { return nil, nil }
var NilContinueFunc ContinueFunc = func(ctx context.Context, session string) (*plugin.RunResponse, error) { return nil, nil }
var NilCacheGroupFunc CacheGroupFunc = func(args string) string { return "" }
var NilClientRendererFunc ClientRendererFunc = func(args string) string { return "" }
