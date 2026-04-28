package tool

import (
	"github.com/manboster/manboster/internal/engine/hook"
	"github.com/manboster/manboster/spec/plugin"
)

type Provider interface {
	plugin.Provider
	RegisterHook(registry *hook.Registry)
}
