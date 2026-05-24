package cron

import (
	"github.com/manboster/manboster/internal/tool"
)

func init() {
	fa := tool.NewFactory[NameType, *Service]()
	fa.RegisterProvider(&Service{})
	fa.RegisterNamespace(NameSet, &runSetInfo)
	fa.RegisterNamespace(NameGet, &runGetInfo)
	fa.RegisterNamespace(NameList, &runListInfo)
	fa.RegisterNamespace(NameDelete, &runDeleteInfo)
	fa.Init()
}
