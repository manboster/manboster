package browser

import (
	"github.com/manboster/manboster/internal/tool"
)

func init() {
	tool.Register(metadata.Name, func() tool.Provider {
		return &Service{}
	})
}
