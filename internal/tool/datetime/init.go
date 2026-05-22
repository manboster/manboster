package datetime

import (
	"github.com/manboster/manboster/internal/tool"
)

func init() {
	fa := tool.NewFactory[NameType, *Service]()
	fa.RegisterProvider(&Service{})
	fa.RegisterNamespace(NameTime, &runTimeInfo)
	fa.RegisterNamespace(NameDate, &runDateInfo)
	fa.Init()
}
