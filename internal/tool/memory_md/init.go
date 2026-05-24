package memory_md

import (
	"github.com/manboster/manboster/internal/tool"
)

func init() {
	fa := tool.NewFactory[NameType, *Service]()
	fa.RegisterProvider(&Service{})
	fa.RegisterNamespace(NameGet, &runGetInfo)
	fa.RegisterNamespace(NameSet, &runSetInfo)
	fa.Init()
}
