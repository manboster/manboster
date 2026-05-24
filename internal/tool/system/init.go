package system

import (
	"github.com/manboster/manboster/internal/tool"
)

func init() {
	fa := tool.NewFactory[NameType, *Service]()
	fa.RegisterProvider(&Service{})
	fa.RegisterNamespace(NameOSInfo, &runOSInfoInfo)
	fa.RegisterNamespace(NameProcessList, &runProcessListInfo)
	fa.RegisterNamespace(NameProcessInfo, &runProcessInfoInfo)
	fa.RegisterNamespace(NameProcessKill, &runProcessKillInfo)
	fa.Init()
}
