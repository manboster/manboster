package datetime

import (
	"github.com/manboster/manboster/internal/tool"
)

func init() {
	if tool.IsLoading {
		fa := tool.NewFactory[NameType, *Service]()
		fa.RegisterProvider(&Service{})
		fa.RegisterNamespace(NameTime, &runTimeInfo)
		fa.RegisterNamespace(NameDate, &runDateInfo)
		fa.Init()
	} else {
		tool.Register(metadata.Name, func() tool.Provider {
			return &Service{}
		})
	}
}
